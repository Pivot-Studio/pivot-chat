package api

//import "C"
import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/service"

	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	//设置ws的超时时间
	wsTimeout = 12 * time.Minute
)

type WsConnContext struct {
	Conn     *websocket.Conn
	UserId   int64
	DeviceId int64
	AppId    int64
}
type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	//DeviceId int64  `json:"device_id"`
	//AppId    int64  `json:"appid"`
}

const (
	PackageType_PT_ERR       PackageType = 0
	PackageType_PT_UNKNOWN   PackageType = 0
	PackageType_PT_SIGN_IN   PackageType = 1
	PackageType_PT_SYNC      PackageType = 2
	PackageType_PT_HEARTBEAT PackageType = 3
	PackageType_PT_MESSAGE   PackageType = 4
	PackageType_PT_JOINGROUP PackageType = 5
)

type (
	PackageType int
	Package     struct {
		Type PackageType `json:"type"`
		Data Input       `json:"data"`
	}
	Input struct {
		model.GroupMessageInput
		model.GroupMessageSyncInput
		model.UserJoinGroupInput
	}
)

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
		logrus.Errorf("[wsHandler] ws upgrade fail, %+v", err)
		return
	}
	// 通过email获取userid
	err = dao.RS.GetUserByEmail(user, user.Email)
	if err != nil {
		logrus.Errorf("[wsHandler] GetUserByEmail fail, %+v", err)
		return
	}
	// conn加入map
	conn.UserId = user.UserId

	// 判断一个用户是否还有别的设备，如果有则下线
	preConn := service.GetConn(user.UserId)
	if preConn != nil {
		preConn.Send("有别的设备登录了你的用户，你寄了", service.PackageType_PT_ERR)
		service.DeleteConn(user.UserId)
		logrus.Info("[wsHandler] Get another conn in same userid-%d, delete pre conn", user.UserId)
	}
	service.SetConn(user.UserId, &conn)
	defer service.DeleteConn(user.UserId) // 出现差错就从map里删除

	err = conn.Send("ws success!waiting for package...", service.PackageType(PackageType_PT_SIGN_IN))
	if err != nil {
		logrus.Errorf("[wsHandler] Send login ack failed, %+v", err)
		return
	}

	//处理连接
	for {
		err = conn.WS.SetReadDeadline(time.Now().Add(wsTimeout))
		_, data, err := conn.WS.ReadMessage()
		if err != nil {
			logrus.Errorf("[wsHandler] ReadMessage failed, %+v", err)
			return
		}
		HandlePackage(data, &conn)
	}
}

// HandlePackage 分类型处理数据包
func HandlePackage(bytes []byte, conn *service.Conn) {
	input := Package{}
	err := json.Unmarshal(bytes, &input)
	if err != nil {
		logrus.Errorf("[HandlePackage] json unmarshal %+v", err)
		//TODO: release连接
		conn.Send(err.Error(), service.PackageType(PackageType_PT_ERR))
		return
	}
	fmt.Printf("%+v\n", input)
	//分类型处理
	//TODO
	switch input.Type {
	case PackageType_PT_UNKNOWN:
		fmt.Println("UNKNOWN")
	case PackageType_PT_SIGN_IN:
		fmt.Println("SIGN_IN")
	case PackageType_PT_SYNC:
		fmt.Println("SYNC")
		err = Sync(input.Data.GroupMessageSyncInput, conn.UserId)
	case PackageType_PT_HEARTBEAT:
		fmt.Println("HEARTBEAT")
	case PackageType_PT_MESSAGE:
		fmt.Println("MESSAGE")
		err = Message(input.Data.GroupMessageInput, conn.UserId)
	case PackageType_PT_JOINGROUP:
		fmt.Println("JOINGROUP")
		err = UserJoinGroup(input.Data.UserJoinGroupInput, conn.UserId)
	default:
		logrus.Info("SWITCH OTHER")
	}
	if err != nil {
		fmt.Println(err)
		conn.Send(err.Error(), service.PackageType(PackageType_PT_ERR))
		return
	}
}

func Message(data model.GroupMessageInput, userId int64) error {
	data.UserId = userId
	fmt.Printf("%+v\n", data)
	return HandleGroupMessage(&data)
}

func Sync(data model.GroupMessageSyncInput, userId int64) error {
	data.UserId = userId
	fmt.Printf("%+v\n", data)
	return HandleSync(&data)
}

func UserJoinGroup(data model.UserJoinGroupInput, userId int64) error {
	data.UserId = userId
	fmt.Printf("%+v\n", data)
	return HandleJoinGroup(&data)
}
