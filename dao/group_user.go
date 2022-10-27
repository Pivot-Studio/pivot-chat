package dao

import "github.com/Pivot-Studio/pivot-chat/model"

func (rs *RdbService) CreateGroupUser(user []*model.GroupUser) error {
	return rs.tx.Create(&user).Error
}
