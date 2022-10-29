package model

import "time"

type GroupUser struct {
	GroupUserId int64 `gorm:"primarykey"` // 自增主键
	GroupId     int64
	UserId      int64     // 用户id
	MemberType  int       // 用户在当前群组的role
	Status      int       // 状态
	CreateTime  time.Time // 创建时间
	UpdateTime  time.Time // 更新时间
}

const (
	OWNER   = 1
	ADMAIN  = 2
	SPEAKER = 3
)

type UserJoinGroupInput struct {
	UserId  int64 `json:"user_id5"` // 发送人userid
	GroupId int64 `json:"group_id5"`
}

type UserJoinGroupOutput struct {
	UserId       int64     `json:"user_id"`
	GroupId      int64     `json:"group_id"`     // 群组id
	Name         string    `json:"name"`         // 组名
	Introduction string    `json:"introduction"` // 群简介
	UserNum      int32     `json:"user_num"`     // 群组人数
	CreateTime   time.Time `json:"create_time"`  // 创建时间
}
