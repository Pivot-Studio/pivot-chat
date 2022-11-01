package model

import (
	"database/sql"
	"time"
)

type User struct {
	UserId         int64 `gorm:"primarykey"`
	UserName       string
	Password       string
	Email          string
	InvitationCode string
	CreateAt       time.Time
	DeleteAt       sql.NullTime
	UpdateAt       time.Time
}

type GetMyGroupResp struct {
	GroupId      int64     `json:"group_id"`
	OwnerId      int64     `json:"owner_id"`
	Name         string    `json:"name"`
	Introduction string    `json:"introduction"`
	UserNum      int32     `json:"user_num"`
	CreateTime   time.Time `json:"create_time"`
}
