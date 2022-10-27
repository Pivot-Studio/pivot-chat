package service

import (
	"errors"
	"sync"
)

var ConnsManager = sync.Map{} // (userID, conn)

func SendToUser(userID int64, data []byte, infoType PackageType) error {
	conn := GetConn(userID)
	if conn == nil {
		return errors.New("[Service] | conn-manager get connection err")
	}
	err := conn.Send(data, infoType)
	return err
}

// SetConn 存储
func SetConn(userID int64, conn *Conn) {
	ConnsManager.Store(userID, conn)
}

// GetConn 获取
func GetConn(userID int64) *Conn {
	value, ok := ConnsManager.Load(userID)
	if ok {
		return value.(*Conn)
	}
	return nil
}

// DeleteConn 删除
func DeleteConn(userID int64) {
	ConnsManager.Delete(userID)
}
