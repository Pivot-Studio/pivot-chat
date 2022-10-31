package service

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Pivot-Studio/pivot-chat/conf"
	"github.com/Pivot-Studio/pivot-chat/constant"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	CHAT_TOKEN_PREFIX = "CHAT_TOKEN_PREFIX"
)
var TokenMap sync.Map

func AddToken(token string, email string)  {
	TokenMap.Store(email, token)
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

func GetUserFromAuth(ctx *gin.Context) (user *model.User, tokenString string, err error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	s := strings.Fields(authHeader)
	if len(s) != 2 || s[0] != "Bearer" {
		return nil, "", constant.TokenLayoutErr
	}
	tokenString = s[1]
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(conf.C.TokenSecret), err
	})
	if err != nil {
		return nil, "", err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, "", constant.UnLoginErr
	}
	email, ok0 := claim["email"].(string)
	uid, ok1 := claim["id"].(float64)
	if !ok0 || !ok1 {
		return nil, "", constant.UnLoginErr
	}

	user = &model.User{
		UserId: int64(uid),
		Email:  email,
	}
	valid := JudgeToken(tokenString, email)
	if !valid {
		return nil, "", constant.TokenLayoutErr
	}
	return user, tokenString, nil
}
