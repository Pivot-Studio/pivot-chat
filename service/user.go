package service

import (
	"context"
	"errors"
	"time"

	"github.com/Pivot-Studio/pivot-chat/util"
	"gorm.io/gorm"

	"github.com/Pivot-Studio/pivot-chat/constant"
	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GetMyGroupResp struct {
	GroupId      int64     `json:"group_id"`
	OwnerId      int64     `json:"owner_id"`
	Name         string    `json:"name"`
	Introduction string    `json:"introduction"`
	UserNum      int32     `json:"user_num"`
	CreateTime   time.Time `json:"create_time"`
}

func Login(email string, password string) (token string, err error) {
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
	// AddToken(token, user.Email)
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
func Register(user *model.User, captcha string) (err error) {
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

func FindUserById(ctx *gin.Context, userid int64) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})
	_, err = GetUserFromAuth(ctx)
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	user.UserId = userid
	err = dao.RS.GetUserbyId(user)
	if err != nil {
		logrus.Fatalf("[Service.FindUserById] FindUserById %+v", err)
		return nil, err
	}
	data["username"] = user.UserName
	data["user_id"] = user.UserId
	data["email"] = user.Email
	return data, nil
}

func ChgPwd(email string, oldPwd string, newPwd string) error {
	return dao.RS.ChangeUserPwd(email, oldPwd, newPwd)
}

func GetMyGroups(UserId int64) (*[]GetMyGroupResp, error) {
	RawGroups, err := dao.RS.GetMyGroups(UserId)
	if err != nil {
		logrus.Errorf("[service.GetMyGroups] %+v", err)
		return nil, err
	}

	var groups []GetMyGroupResp
	for _, r := range RawGroups {
		groups = append(groups, GetMyGroupResp{
			GroupId:      r.GroupId,
			OwnerId:      r.OwnerId,
			Name:         r.Name,
			Introduction: r.Introduction,
			UserNum:      r.UserNum,
			CreateTime:   r.CreateTime,
		})
	}

	return &groups, nil
}
