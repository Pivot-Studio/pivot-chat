package dao

import (
	"errors"

	"github.com/Pivot-Studio/pivot-chat/model"
	"gorm.io/gorm"
)

func (rs *RdbService) CreateUser(user []*model.User) error {
	return rs.tx.Create(&user).Error
}

func (rs *RdbService) UserExist(email string) (exist bool, err error) {
	err = rs.tx.Where("email = ?", email).First(&model.User{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}
func (rs *RdbService) GetUser(user *model.User) (err error) {
	err = rs.tx.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func (rs *RdbService) UpdateUser(user *model.User) (err error) {
	err = rs.tx.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (rs *RdbService) GetUserbyId(user *model.User) (err error) {
	err = rs.tx.Where("user_id = ?", user.UserId).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}