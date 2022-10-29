package api

//import "C"
import (
	"encoding/json"
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

// type LoginResponse struct {
// 	Msg  string `json:"msg"`
// 	data struct {
// 		Username string `json:"username"`
// 		UserId   int64  `json:"user_id"`
// 		Email    string `json:"email"`
// 	}
// }

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
	//todo 紧急解决
	req := LoginInfo{
		Email:    ctx.Query("email"),
		Password: ctx.Query("password"),
	}
	logrus.Infof("email:%s password:%s", req.Email, req.Password)
	if req.Email == "" || req.Password == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"msg": "登录失败, 账号密码不能为空",
		})
		return
	}

	if !service.Auth(req.Email, req.Password) {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"msg": "登录失败, 账号密码错误或不匹配",
		})
		return
	}
	var err error
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
	user := model.User{}
	err = dao.RS.GetUserByEmail(&user, req.Email)
	if err != nil {
		logrus.Errorf("[wsHandler] GetUserByEmail fail, %+v", err)
		return
	}
	// conn加入map
	conn.UserId = user.UserId
	service.SetConn(user.UserId, &conn)

	err = conn.Send("login success! waiting for package...", service.PackageType(PackageType_PT_SIGN_IN))
	if err != nil {
		logrus.Errorf("[wsHandler] Send login ack failed, %+v", err)
		service.DeleteConn(user.UserId) // 出现差错就从map里删除
		return
	}

	//处理连接
	for {
		err = conn.WS.SetReadDeadline(time.Now().Add(wsTimeout))
		_, data, err := conn.WS.ReadMessage()
		if err != nil {
			logrus.Errorf("[wsHandler] ReadMessage failed, %+v", err)
			service.DeleteConn(user.UserId) // 出现差错就从map里删除
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
	// err := json.Unmarshal(data, &meg)
	// if err != nil {
	// 	logrus.Errorf("[Message] json unmarshal %+v", err)
	// 	return err
	// }
	// meg.UserId = userId
	return HandleGroupMessage(&data)
}

func Sync(data model.GroupMessageSyncInput, userId int64) error {
	data.UserId = userId
	// err := json.Unmarshal([]byte(data), &meg)
	// if err != nil {
	// 	logrus.Errorf("[Message] json unmarshal %+v", err)
	// 	return err
	// }
	// meg.UserId = userId
	return HandleSync(&data)
}

func UserJoinGroup(data model.UserJoinGroupInput, userId int64) error {
	data.UserId = userId
	// err := json.Unmarshal(data, &meg)
	// if err != nil {
	// 	logrus.Errorf("[Message] json unmarshal %+v", err)
	// 	return err
	// }
	// meg.UserId = userId
	return HandleJoinGroup(&data)
}
