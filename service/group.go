package service

import (
	"time"
)

// Group 群组
type Group struct {
	Id           int64   `gorm:"primarykey"`    // 群组id
	Name         string      // 组名
	Introduction string      // 群简介
	UserNum      int32       // 群组人数
	CreateTime   time.Time   // 创建时间
	UpdateTime   time.Time   // 更新时间
	Members      []GroupUser `gorm:"-"` // 群组成员
}

type GroupUser struct {
	Id         int64   `gorm:"primarykey"`  // 自增主键
	UserId     int64     // 用户id
	MemberType int       // 用户在当前群组的role
	Status     int       // 状态
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间
}

type SendInfo struct {
	UserId     int64     // 用户id
	Messgae    string    // 消息内容
}

func(g *Group) SendMessgae(sendInfo SendInfo) {
	
}