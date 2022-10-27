package service

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Conn struct {
	WSMutex sync.Mutex      // WS写锁
	WS      *websocket.Conn // websocket连接
	UserId  int64           // 用户ID
}

type PackageType int
type Package struct {
	//数据包内容, 按需修改
	Type PackageType
	Data []byte
}

const (
	PackageType_PT_ERR   PackageType = 0
	PackageType_PT_UNKNOWN   PackageType = 0
	PackageType_PT_SIGN_IN   PackageType = 1
	PackageType_PT_SYNC      PackageType = 2
	PackageType_PT_HEARTBEAT PackageType = 3
	PackageType_PT_MESSAGE   PackageType = 4
)

func (c *Conn) Send(data []byte, t PackageType) error {
	c.WSMutex.Lock()
	defer c.WSMutex.Unlock()
	ret := Package{
		Type: t,
		Data: data,
	}
	err := c.WS.SetWriteDeadline(time.Now().Add(10 * time.Millisecond))
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(ret)
	if err != nil {
		return err
	}
	return c.WS.WriteMessage(websocket.BinaryMessage, bytes)
}
