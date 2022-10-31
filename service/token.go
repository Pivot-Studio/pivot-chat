package service

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Pivot-Studio/pivot-chat/conf"
	"github.com/Pivot-Studio/pivot-chat/constant"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

const (
	CHAT_TOKEN_PREFIX = "CHAT_TOKEN_PREFIX"
)

var TokenMap sync.Map

func AddToken(token string, email string) {
	TokenMap.Store(email, token)
}

func GetToken(email string) string {
	value, ok := TokenMap.Load(email)
	if !ok {
		fmt.Println("[Service.GetToken] map load failed")
		return ""
	}
	return value.(string)
}

func JudgeToken(token string, email string) bool {
	value, ok := TokenMap.Load(email)
	if !ok {
		fmt.Println("[Service.GetToken] map load failed")
		return false
	}
	if value.(string) != token {
		return false
	}
	return true
}
func DeleteToken(email string) {
	_, ok := TokenMap.LoadAndDelete(email)
	if !ok {
		fmt.Println("[Service.DeleteToken] map LoadAndDelete failed")
	}
}

func ParseToken(tokenString string) (email string, id int64, tokenTime time.Time, err error) {
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(conf.C.TokenSecret), err
	})
	if err != nil {
		return "", 0, time.Time{}, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", 0, time.Time{}, constant.UnLoginErr
	}
	email, ok0 := claim["email"].(string)
	uid, ok1 := claim["id"].(float64)
	timeStamp, ok2 := claim["timeStamp"].(string)
	if !ok0 || !ok1 || !ok2 {
		return "", 0, time.Time{}, constant.UnLoginErr
	}
	tokenTime, err = time.ParseInLocation("2006-01-02 15:04:05", timeStamp, time.Local)
	if err != nil || tokenTime.IsZero() {
		return "", 0, time.Time{}, constant.UnLoginErr
	}
	return email, int64(uid), tokenTime, nil
}

func WSLoginAuth(tokenString string) (user *model.User, err error) {
	curTokenemail, curTokenid, curTokenTime, err := ParseToken(tokenString)
	if err != nil {
		logrus.Errorf("[Service.WSLoginAuth] ParseToken err:%+v", err)
		return nil, err
	}
	preToken := GetToken(curTokenemail)
	if preToken == "" {
		user = &model.User{
			UserId: int64(curTokenid),
			Email:  curTokenemail,
		}
		return user, nil
	}
	_, _, preTokenTime, err := ParseToken(preToken)
	if err != nil {
		logrus.Errorf("[Service.WSLoginAuth] ParseToken err:%+v", err)
		return nil, err
	}
	if preTokenTime.After(curTokenTime) {
		return nil, errors.New("登录失败，存在更新的token")
	}

	user = &model.User{
		UserId: int64(curTokenid),
		Email:  curTokenemail,
	}
	return user, nil
}
func GetUserFromAuth(ctx *gin.Context) (user *model.User, err error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	s := strings.Fields(authHeader)
	if len(s) != 2 || s[0] != "Bearer" {
		return nil, constant.TokenLayoutErr
	}
	tokenString := s[1]
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(conf.C.TokenSecret), err
	})
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, constant.UnLoginErr
	}
	email, ok0 := claim["email"].(string)
	uid, ok1 := claim["id"].(float64)
	if !ok0 || !ok1 {
		return nil,  constant.UnLoginErr
	}

	user = &model.User{
		UserId: int64(uid),
		Email:  email,
	}
	valid := JudgeToken(tokenString, email)
	if !valid {
		return nil,  constant.TokenLayoutErr
	}
	return user,  nil
}
