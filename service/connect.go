package service

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Conn struct {
	WSMutex  sync.Mutex      // WS写锁
	WS       *websocket.Conn // websocket连接
	UserId   int64           // 用户ID
}