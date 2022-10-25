package model

import "time"

type Message struct {
	ID           int64     `gorm:"primarykey"` // 自增主键
	SenderType   int32     // 发送者类型
	SenderID     int64     // 发送者账户id
	ReceiverType int32     // 接收者账户id
	ReceiverID   int64     `gorm:"unique_index:u_meg"` // 接收者id,如果是单聊信息，则为user_id，如果是群组消息，则为group_id
	Content      []byte    // 消息内容
	Seq          int64     `gorm:"unique_index:u_meg"` // 消息同步序列
	SendTime     time.Time // 消息发送时间
}
