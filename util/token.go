package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"pivot-chat/conf"
	"pivot-chat/model"
)

func GenerateToken(u *model.User) (token string, err error) {
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        u.UserId,
		"email":     u.Email,
		"timeStamp": time.Now().Format("2006-01-02 15:04:05"),
	}).SignedString([]byte(conf.C.TokenSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}
