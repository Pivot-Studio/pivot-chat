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

type GroupMessageSyncInput struct {
	UserId     int64 // 发送人userid
	GroupId int64 // 群组id
	SyncSeq int64 // 开始同步的seq，是用户的本地seq+1
}

type GroupMessageSyncOutput struct {
	UserId     int64 // 接受者user_id
	GroupId    int64 // 群组id
	Data       []GroupMessageOutput
	MaxSeq     int64  // 该群组最新Seq
}