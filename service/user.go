package service

import (
	"context"
	"github.com/Pivot-Studio/pivot-chat/constant"
	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, user *model.User, captcha string) (err error) {
	res, err := dao.Cache.Get(context.Background(), user.Email).Result()
	if err != nil {
		return err
	}
	if res != captcha {
		return constant.CaptchaErr
	}
	exist, err := dao.RS.UserExist(user.Email)
	if err != nil {
		return err
	}
	if exist {
		return constant.EmailExistErr
	}
	err = dao.RS.CreateUser([]*model.User{user})
	if err != nil {
		return err
	}
	return nil
}