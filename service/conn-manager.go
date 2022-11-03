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

func walk(key, value interface{}) bool {
	logrus.Debug("Key =", key, "Value =", value)
	return true
}

// SetConn 存储
func SetConn(userID int64, conn *Conn) {
	logrus.Debug("Before SetConn")
	ConnsManager.Range(walk)
	ConnsManager.Store(userID, conn)
	logrus.Debug("After SetConn")
	ConnsManager.Range(walk)
}

// GetConn 获取
func GetConn(userID int64) *Conn {
	logrus.Debug("GetConn")
	ConnsManager.Range(walk)
	value, ok := ConnsManager.Load(userID)
	if ok {
		return value.(*Conn)
	}
	return nil
}

// DeleteConn 删除
func DeleteConn(userID int64) {
	logrus.Debug("Before DeleteConn")
	ConnsManager.Range(walk)
	value, ok := ConnsManager.LoadAndDelete(userID)
	if ok {
		err := value.(*Conn).WS.Close()
		if err != nil {
			logrus.Errorf("delete user-%d err:%+v", userID, err)
		}
		logrus.Info("delete user:", userID, " Conn!")
	}
	logrus.Debug("After DeleteConn")
	ConnsManager.Range(walk)
	return
}
