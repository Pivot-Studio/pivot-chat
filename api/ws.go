package api

import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	//设置ws的超时时间
	wsTimeout = 12 * time.Minute
)

type PackageType int
type Package struct {
	//数据包内容, 按需修改
	Type PackageType
	Id   int64
	data []byte
}
type WsConnContext struct {
	Conn     *websocket.Conn
	UserId   int64
	DeviceId int64
	AppId    int64
}

const (
	PackageType_PT_UNKNOWN   PackageType = 0
	PackageType_PT_SIGN_IN   PackageType = 1
	PackageType_PT_SYNC      PackageType = 2
	PackageType_PT_HEARTBEAT PackageType = 3
	PackageType_PT_MESSAGE   PackageType = 4
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(ctx *gin.Context) {
	//TODO:auth 这里鉴权, 成功就修改一下下面wsConn的id
	c := WsConnContext{}
	var err error

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
	case PackageType_PT_HEARTBEAT:
		fmt.Println("HEARTBEAT")
	case PackageType_PT_MESSAGE:
		fmt.Println("MESSAGE")
	default:
		logrus.Info("SWITCH OTHER")
	}
}