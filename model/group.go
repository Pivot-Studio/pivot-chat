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
	Members      []*GroupUser `gorm:"-"` // 群组成员
}

type GroupUser struct {
	GroupUserId int64     `gorm:"primarykey"` // 自增主键
	GroupId     int64     `gorm:"unique_index:u_group"`
	UserId      int64     `gorm:"unique_index:u_group"` // 用户id
	MemberType  int       // 用户在当前群组的role
	Status      int       // 状态
	CreateTime  time.Time // 创建时间
	UpdateTime  time.Time // 更新时间
}
