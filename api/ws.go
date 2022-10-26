package api

//import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/Pivot-Studio/pivot-chat/service"
	"net/http"
	"time"

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
	DeviceId int64
	AppId    int64
}
type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	DeviceId int64  `json:"device_id"`
	AppId    int64  `json:"appid"`
}

const (
	PackageType_PT_UNKNOWN   PackageType = 0
	PackageType_PT_SIGN_IN   PackageType = 1
	PackageType_PT_SYNC      PackageType = 2
	PackageType_PT_HEARTBEAT PackageType = 3
	PackageType_PT_MESSAGE   PackageType = 4
)

type PackageType int
type Package struct {
	//数据包内容, 按需修改
	Type PackageType
	Data []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(ctx *gin.Context) {
	req := LoginInfo{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logrus.Fatalf("[api.wsHandler] BindJson %+v", err)
	}
	if !service.Auth(req.Email, req.Password) {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"msg": "登录失败, 账号密码错误或不匹配",
		})
		return
	}

	//登录成功, 升级为websocket
	c := WsConnContext{
		AppId:    req.AppId,
		DeviceId: req.DeviceId,
	}
	c.Conn, err = upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Errorf("[wsHandler] ws upgrade fail, %+v", err)
	}

	//处理连接
	for {
		err = c.Conn.SetReadDeadline(time.Now().Add(wsTimeout))
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			logrus.Errorf("[wsHandler] ReadMessage failed, %+v", err)
			return
		}
		c.HandlePackage(data)
	}
}

// HandlePackage 分类型处理数据包
func (c *WsConnContext) HandlePackage(bytes []byte) {
	input := Package{}
	err := json.Unmarshal(bytes, &input)
	if err != nil {
		logrus.Errorf("[HandlePackage] json unmarshal %+v", err)
		//TODO: release连接
		return
	}

	//分类型处理
	//TODO
	switch input.Type {
	case PackageType_PT_UNKNOWN:
		fmt.Println("UNKNOWN")
	case PackageType_PT_SIGN_IN:
		fmt.Println("SIGN_IN")
	case PackageType_PT_SYNC:
		fmt.Println("SYNC")
		c.Sync(input.Data)
	case PackageType_PT_HEARTBEAT:
		fmt.Println("HEARTBEAT")
	case PackageType_PT_MESSAGE:
		fmt.Println("MESSAGE")
		c.Message(input.Data)
	default:
		logrus.Info("SWITCH OTHER")
	}
}

func (c *WsConnContext) Message(data []byte) {
	meg := model.GroupMessageInput{}
	err := json.Unmarshal(data, &meg)
	if err != nil {
		logrus.Errorf("[Message] json unmarshal %+v", err)
		return
	}
	HandleGroupMessage(&meg)
}

func (c *WsConnContext) Sync(data []byte) {
	meg := model.GroupMessageSyncInput{}
	err := json.Unmarshal(data, &meg)
	if err != nil {
		logrus.Errorf("[Message] json unmarshal %+v", err)
		return
	}
	HandleSync(&meg)
}
