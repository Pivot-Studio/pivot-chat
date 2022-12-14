package util

import (
	"github.com/Pivot-Studio/pivot-chat/constant"
	"golang.org/x/crypto/bcrypt"
)

func EncodePassword(password string) (passwordHash string, err error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", constant.HashErr
	}
	return string(b), nil
}

func ComparePassword(passwordInDB string, password string) (match bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(passwordInDB), []byte(password)); err != nil {
		return false
	}
	return true
}
