package service

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Conn struct {
	WSMutex sync.Mutex      // WS写锁
	WS      *websocket.Conn // websocket连接
	UserId  int64           // 用户ID
}

func (c *Conn) Send(bytes []byte) error {
	c.WSMutex.Lock()
	defer c.WSMutex.Unlock()

	err := c.WS.SetWriteDeadline(time.Now().Add(10 * time.Millisecond))
	if err != nil {
		return err
	}
	return c.WS.WriteMessage(websocket.BinaryMessage, bytes)
}
