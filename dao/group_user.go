package dao

import (
	"pivot-chat/model"
)

func (rs *RdbService) CreateGroupUser(user []*model.GroupUser) error {
	return rs.tx.Create(&user).Error
}

func (rs *RdbService) GetGroupUsers(groupID int64) (*[]model.GroupUser, error) {
	var g []model.GroupUser
	err := rs.tx.Table("group_users").Where("group_id = ?", groupID).Find(&g).Error
	return &g, err
}
