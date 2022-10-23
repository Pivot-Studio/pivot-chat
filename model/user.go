package model

import (
	"database/sql"
	"time"
)

type User struct {
	UserId    int64 `gorm:"primarykey"`
	UserName  string
	Password  string
	Email     string
	InvitationCode string
	CreateAt  time.Time
	DeleteAt  sql.NullTime
	UpdateAt  time.Time
}