package service

import (
	"fmt"
	"time"

	"github.com/Pivot-Studio/pivot-chat/constant"
	"github.com/sirupsen/logrus"
)

// Group 群组
type Group struct {
	Id           int64       `gorm:"primarykey"` // 群组id
	Name         string      // 组名
	Introduction string      // 群简介
	UserNum      int32       // 群组人数
	CreateTime   time.Time   // 创建时间
	UpdateTime   time.Time   // 更新时间
	Members      []GroupUser `gorm:"-"` // 群组成员
}

type GroupUser struct {
	Id         int64     `gorm:"primarykey"` // 自增主键
	UserId     int64     // 用户id
	MemberType int       // 用户在当前群组的role
	Status     int       // 状态
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间
}

type SendInfo struct {
	UserId     int64  // 用户id
	Messgae    string // 消息内容
	SenderType int    // 发送者身份
}

const (
	SenderType_USER = 1
)

func (g *Group) IsMember(userId int64) bool {
	for i := range g.Members {
		if g.Members[i].UserId == userId {
			return true
		}
	}
	return false
}

func (g *Group) SendMessgae(sendInfo SendInfo) error {
	if sendInfo.SenderType == SenderType_USER && !g.IsMember(sendInfo.UserId) {
		logrus.Fatalf("[Service] | group sendmeg error: user isn't in group | sendInfo:", sendInfo)
		return constant.UserNotMatchGroup
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered. Error:\n", r)
			}
		}()
		// 将消息发送给群组用户
		for _, user := range g.Members {
			// 前面已经发送过，这里不需要再发送
			if sendInfo.SenderType == SenderType_USER && user.UserId == sendInfo.UserId {
				continue
			}
			// TODO:发送消息
		}
	}()
	return nil
}
