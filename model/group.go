package model

import "time"

// Group 群组
type Group struct {
	GroupId      int64     `gorm:"primarykey"` // 群组id
	Name         string    // 组名
	Introduction string    // 群简介
	UserNum      int32     // 群组人数
	CreateTime   time.Time // 创建时间
	UpdateTime   time.Time // 更新时间
	MaxSeq       int64
}

type GroupUser struct {
	GroupUserId int64 `gorm:"primarykey"` // 自增主键
	UserId      int64 // 用户id
	GroupID     int64
	MemberType  int       // 用户在当前群组的role
	Status      int       // 状态
	CreateTime  time.Time // 创建时间
	UpdateTime  time.Time // 更新时间
}

type GroupMessageInput struct {
	UserId  int64 // 发送人userid
	GroupId int64 // 群组id
	Data    string
}

type GroupMessageOutput struct {
	UserId   int64 // 接受者user_id
	GroupId  int64 // 群组id
	Data     string
	SenderId int64 // 发送者账户id
	Seq      int64 // 该条消息的正确seq
}
