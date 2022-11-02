package service

import (
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
)

var ConnsManager = sync.Map{} // (userID, conn)

func SendToUser(userID int64, data interface{}, infoType PackageType) error {
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
	value, ok := ConnsManager.LoadAndDelete(userID)
	if ok {
		err := value.(*Conn).WS.Close()
		if err != nil {
			logrus.Errorf("delete user-%d err:%+v", userID, err)
		}
		logrus.Info("delete user:", userID, " Conn!")
	}
	return
}
