package api

//import "C"
import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/proto/github.com/Pivot-Studio/pivot-chat/pb"
	"github.com/Pivot-Studio/pivot-chat/service"
	"google.golang.org/protobuf/proto"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	//设置ws的超时时间
	wsTimeout = 5 * time.Minute
)

type WsConnContext struct {
	Conn     *websocket.Conn
	UserId   int64
	DeviceId int64
	AppId    int64
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(ctx *gin.Context) {
	token := ctx.Query("token")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"msg":  "登录失败",
			"data": errors.New("token缺失"),
		})
		return
	}
	user, err := service.WSLoginAuth(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"msg":  "登录失败",
			"data": err.Error(),
		})
		return
	}
	service.AddToken(token, user.Email)
	defer service.DeleteToken(user.Email)
	// 登录成功, 升级为websocket
	conn := service.Conn{
		WSMutex: sync.Mutex{},
	}
	conn.WS, err = upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Errorf("[ws-Handler] ws upgrade fail, %+v", err)
		return
	}
	// 通过email获取userid
	err = dao.RS.GetUserByEmail(user, user.Email)
	if err != nil {
		logrus.Errorf("[ws-Handler] GetUserByEmail fail, %+v", err)
		return
	}
	// conn加入map
	conn.UserId = user.UserId

	// 判断一个用户是否还有别的设备，如果有则下线
	preConn := service.GetConn(user.UserId)
	if preConn != nil {
		preConn.Send(&pb.ErrorMessage{
			Code:    1,
			Message: "有别的用户登录，你寄楽!",
		}, pb.PackageType_ERR)
		service.DeleteConn(user.UserId)
		logrus.Info("[ws-Handler] Get another conn in same userid-", user.UserId, ", delete pre conn")
	}
	service.SetConn(user.UserId, &conn)
	defer service.DeleteConn(user.UserId) // 出现差错就从map里删除

	err = conn.Send(&pb.ErrorMessage{
		Code:    0,
		Message: "---signin success! waiting for package...---",
	}, pb.PackageType_SIGHIN)
	if err != nil {
		logrus.Errorf("[ws-Handler] Send login ack failed, %+v", err)
		return
	}

	//处理连接
	for {
		err = conn.WS.SetReadDeadline(time.Now().Add(wsTimeout))
		if err != nil {
			logrus.Errorf("[ws-Handler] SetReadDeadline failed, %+v", err)
			return
		}
		_, data, err := conn.WS.ReadMessage()
		if err != nil {
			logrus.Errorf("[ws-Handler] ReadMessage failed, %+v", err)
			return
		}
		HandlePackage(data, &conn)
	}
}

// HandlePackage 分类型处理数据包
func HandlePackage(bytes []byte, conn *service.Conn) {
	//分类型处理
	//TODO
	pkg := &pb.Input{}
	err := proto.Unmarshal(bytes, pkg)
	if err != nil {
		logrus.Errorf("[HandlePackage] proto unmarshal %+v", err)
		//TODO: release连接
		conn.Send(&pb.ErrorMessage{
			Code:    1,
			Message: err.Error(),
		}, pb.PackageType_ERR)
		return
	}
	switch pkg.Type {
	case pb.PackageType_ERR:
		fmt.Println("UNKNOWN")
	case pb.PackageType_SIGHIN:
		fmt.Println("SIGN_IN")
	case pb.PackageType_MESSGAE:
		fmt.Println("MESSAGE")
		err = Message(pkg.GetData(), conn.UserId)
	case pb.PackageType_JOINGROUP:
		fmt.Println("JOINGROUP")
		err = UserJoinGroup(pkg.GetData(), conn.UserId)
	default:
		logrus.Info("SWITCH OTHER")
	}
	if err != nil {
		conn.Send(&pb.ErrorMessage{
			Code:    1,
			Message: err.Error(),
		}, pb.PackageType_ERR)
		return
	}
}

func Message(data []byte, userId int64) error {
	req := pb.GroupMessageRequest{}
	err := proto.Unmarshal(data, &req)
	if err != nil {
		logrus.Errorf("[ws-Message] proto unmarshal falied:%+v", err)
		return err
	}
	req.UserId = userId
	return HandleGroupMessage(&req)
}

func UserJoinGroup(data []byte, userId int64) error {
	req := pb.UserJoinGroupRequest{}
	err := proto.Unmarshal(data, &req)
	if err != nil {
		logrus.Errorf("[ws-UserJoinGroup] proto unmarshal falied:%+v", err)
		return err
	}
	req.UserId = userId
	return HandleJoinGroup(&req)
}
