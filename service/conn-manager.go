package service

import "sync"

var ConnsManager = sync.Map{}

func SendToUser()

// SetConn 存储
func SetConn(deviceId int64, conn *Conn) {
	ConnsManager.Store(deviceId, conn)
}

// GetConn 获取
func GetConn(deviceId int64) *Conn {
	value, ok := ConnsManager.Load(deviceId)
	if ok {
		return value.(*Conn)
	}
	return nil
}

// DeleteConn 删除
func DeleteConn(deviceId int64) {
	ConnsManager.Delete(deviceId)
}
