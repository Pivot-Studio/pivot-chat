package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

var ConnsManager = sync.Map{} // (userID, conn)

func SendToUser(userID int64, data []byte, infoType PackageType) error {
	conn := GetConn(userID)
	if conn == nil {
		return errors.New(fmt.Sprintf("[Service] | conn-manager get user:%d connection err", userID))
	}
	err := conn.Send(data, infoType)
	return err
}

func walk(key, value interface{}) bool {
	logrus.Info("Key =", key, "Value =", value)
	return true
}

// SetConn 存储
func SetConn(userID int64, conn *Conn) {
	logrus.Info("Before SetConn")
	ConnsManager.Range(walk)
	ConnsManager.Store(userID, conn)
	logrus.Info("After SetConn")
	ConnsManager.Range(walk)
}

// GetConn 获取
func GetConn(userID int64) *Conn {
	logrus.Info("GetConn")
	ConnsManager.Range(walk)
	value, ok := ConnsManager.Load(userID)
	if ok {
		return value.(*Conn)
	}
	return nil
}

// DeleteConn 删除
func DeleteConn(userID int64) {
	logrus.Info("Before DeleteConn")
	ConnsManager.Range(walk)
	value, ok := ConnsManager.LoadAndDelete(userID)
	if ok {
		err := value.(*Conn).WS.Close()
		if err != nil {
			logrus.Errorf("delete user-%d err:%+v", userID, err)
		}
		logrus.Info("delete user:", userID, " Conn!")
	}
	logrus.Info("After DeleteConn")
	ConnsManager.Range(walk)
	return
}
