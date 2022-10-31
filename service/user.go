package service

import (
	"context"
	"errors"

	"github.com/Pivot-Studio/pivot-chat/util"
	"gorm.io/gorm"

	"github.com/Pivot-Studio/pivot-chat/constant"
	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Login(email string, password string) (token string, err error){
	user, valid := auth(email, password)
	if !valid {
		logrus.Errorf("[Service.Login] auth %+v", err)
		return "", constant.UnLoginPwdErr
	}
	token, err = util.GenerateToken(user)
	if err != nil {
		logrus.Errorf("[Service.Login] GenerateToken %+v", err)
		return "", errors.New("生成token失败")
	}
	AddToken(token, user.Email)
	return token, nil
}
func auth(email string, password string) (*model.User, bool) {
	user := &model.User{}
	err := dao.RS.GetUserByEmail(user, email)
	if err != nil {
		logrus.Fatalf("[Service.Auth] GetUserByEmail file %+v", err)
		return nil, false
	}
	return user, util.ComparePassword(user.Password, password)
}
func Register(ctx *gin.Context, user *model.User, captcha string) (err error) {
	//邮箱验证码部分
	codeKey := CHAT_CODE_PREFIX + user.Email
	res, err := dao.Cache.Get(context.Background(), codeKey).Result()
	if err != nil {
		return err
	}
	if res != captcha {
		return constant.CaptchaErr
	}
	err = dao.RS.GetUserByEmail(&model.User{Email: user.Email}, user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	} else if err == nil {
		return errors.New("该邮箱已被注册")
	}
	err = dao.RS.CreateUser([]*model.User{user})
	if err != nil {
		return err
	}
	return nil
}



func ChgPwd(ctx *gin.Context, email string, oldPwd string, newPwd string) error {
	return dao.RS.ChangeUserPwd(email, oldPwd, newPwd)
}
