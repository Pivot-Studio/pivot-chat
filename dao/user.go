package dao

import (
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"pivot-chat/model"
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

func (rs *RdbService) ChangeUserPwd(email string, oldPwd string, newPwd string) (err error) {
	user := model.User{}
	err = rs.GetUserByEmail(&user, email)
	if err != nil {
		return err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPwd)) != nil {
		return errors.New("the password is wrong please try again")
	}

	user.Password = newPwd

	err = rs.UpdateUser(&user)

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
	return rs.tx.Table("users").Where("email = ?", Email).First(user).Error
}

func (rs *RdbService) GetMyGroups(UserId int64) ([]model.Group, error) {
	var groups []model.Group
	err := rs.tx.Table("groups").Where("owner_id = ?", UserId).Find(&groups).Error
	if err != nil {
		logrus.Errorf("[dao.GetMyGroups] %+v", err)
		return nil, err
	}
	return groups, nil
}

func (rs *RdbService) GetMyJoinedGroups(UserId int64) ([]model.Group, error) {
	var group_users []model.GroupUser
	err := rs.tx.Table("group_users").Where("user_id = ?", UserId).Find(&group_users).Error
	if err != nil {
		logrus.Errorf("[dao.GetMyJoinedGroups] %+v", err)
		return nil, err
	}
	ret := make([]model.Group, 0)
	for _, user := range group_users {
		group, err := RS.QueryGroup(user.GroupId)
		if err != nil {
			logrus.Errorf("[dao.GetMyJoinedGroups.QueryGroup] %+v", err)
			continue
		}
		ret = append(ret, *group)
	}
	return ret, nil
}
