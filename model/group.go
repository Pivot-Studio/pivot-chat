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

type GroupMessageInput struct {
	UserId  int64  `json:"user_id4"`  // 发送人userid
	GroupId int64  `json:"group_id4"` // 群组id
	Data    string `json:"data4"`
	ReplyTo int64  `json:"reply_to4"`
	Type    int64  `json:"type4"`
}

type GroupMessageOutput struct {
	UserId   int64     `json:"user_id"`  // 接受者user_id
	GroupId  int64     `json:"group_id"` // 群组id
	Data     string    `json:"data"`
	SenderId int64     `json:"sender_id"` // 发送者账户id
	Seq      int64     `json:"seq"`       // 该条消息的正确seq
	ReplyTo  int64     `json:"reply_to"`
	Type     int64     `json:"type"`
	Time     time.Time `json:"time"`
}
