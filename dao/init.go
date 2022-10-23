package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/Pivot-Studio/pivot-chat/conf"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	RS    *RdbService
	Cache *redis.Client
)

type RdbService struct {
	tx *gorm.DB
}

func init() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4,utf8",
		conf.C.DB.UserName,
		conf.C.DB.Password,
		conf.C.DB.Host,
		conf.C.DB.Schema,
	)
	engine, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("[init]database connect error %v", err)
	}
	DB = engine
	RS = &RdbService{
		DB,
	}
	sqldb, err := engine.DB()
	if err != nil {
		logrus.Fatalf("[init]invalid database driver %v", err)
	}
	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(10000)
	sqldb.SetConnMaxLifetime(time.Second * 3)
	DB.AutoMigrate(model.User{})
	logrus.Info("[init] db init")
}
func init() {
	Cache = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     conf.C.Redis.Host,
		Password: conf.C.Redis.Password,
	})
	res, err := Cache.Ping(context.Background()).Result()
	if err != nil || res != "PONG" {
		logrus.Fatalf("[init] init redis err: %+v", err)
	}
}
