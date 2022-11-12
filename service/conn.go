package service

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"pivot-chat/pkg/pb"
)

type Conn struct {
	WSMutex sync.Mutex      // WS写锁
	WS      *websocket.Conn // websocket连接
	UserId  int64           // 用户ID
}

func (c *Conn) Send(message proto.Message, pkType pb.PackageType) error {
	output := pb.Output{
		Type: pkType,
	}

	if message != nil {
		msgBytes, err := proto.Marshal(message)
		if err != nil {
			return err
		}
		output.Data = msgBytes
	}

	outputBytes, err := proto.Marshal(&output)
	if err != nil {
		return err
	}

	c.WSMutex.Lock()
	defer c.WSMutex.Unlock()
	err = c.WS.SetWriteDeadline(time.Now().Add(200 * time.Millisecond))
	if err != nil {
		return err
	}
	return c.WS.WriteMessage(websocket.TextMessage, outputBytes)
}
