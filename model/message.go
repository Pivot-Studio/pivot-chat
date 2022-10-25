package model

import "time"

type Message struct {
	MessageId           int64     `gorm:"primarykey"` // 自增主键
	SenderType   int     // 发送者类型
	SenderId     int64     // 发送者账户id
	ReceiverType int     // 接收者账户id
	ReceiverId   int64     `gorm:"unique_index:u_meg"` // 接收者id,如果是单聊信息，则为user_id，如果是群组消息，则为group_id
	Content      []byte    // 消息内容
	Seq          int64     `gorm:"unique_index:u_meg"` // 消息同步序列
	SendTime     time.Time // 消息发送时间
}
