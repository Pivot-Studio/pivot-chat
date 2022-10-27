package dao

import (
	"errors"
	"github.com/Pivot-Studio/pivot-chat/model"
)

func (rs *RdbService) CreateUser(user []*model.User) error {

	return rs.tx.Create(&user).Error
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
func (rs *RdbService) GetUserbyUsername(user *model.User) (err error) {
	err = rs.tx.Where("user_name = ?", user.UserName).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (rs *RdbService) ChangeUserPwd(user *model.User, oldPwd string, newPwd string) (err error) {

	err = rs.GetUserbyUsername(user)
	if err != nil {
		return err
	}

	if user.Password != oldPwd {
		return errors.New("the password is wrong please try again")
	}

	user.Password = newPwd

	err = rs.UpdateUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (rs *RdbService) ChangeUserName(user *model.User, newUserName string) (err error) {

	user.UserName = newUserName

	err = rs.UpdateUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (rs *RdbService) GetUserByEmail(user *model.User, Email string) error {
	user = &model.User{}
	return rs.tx.Table("users").Where("email = ?", Email).First(user).Error
}
