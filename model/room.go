package model

import (
	"container/list"
	"sync"
)

type Room struct {
	RoomId int64      // 房间ID
	Conns  *list.List // 订阅房间消息的连接
	lock   sync.RWMutex
}