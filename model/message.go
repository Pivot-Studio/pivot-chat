package model

import "time"

type Message struct {
	MessageId    int64     `gorm:"primarykey"` // 自增主键
	SenderId     int64     // 发送者账户id
	ReceiverId   int64     `gorm:"unique_index:u_meg"` // 接收者id,均为group_id
	Content      string    // 消息内容
	Seq          int64     `gorm:"unique_index:u_meg"` // 消息同步序列
	SendTime     time.Time // 消息发送时间（落库时间）
}

type GroupMessageInput struct {
	UserId     int64 // 发送人userid
	GroupId int64 // 群组id
	Data       string
}

type GroupMessageOutput struct {
	UserId     int64 // 接受者user_id
	GroupId    int64 // 群组id
	Data       string
	SenderId   int64    // 发送者账户id
	Seq        int64  // 该条消息的正确seq
}

