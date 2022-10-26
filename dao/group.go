package dao

import (
	"github.com/Pivot-Studio/pivot-chat/model"
	"gorm.io/gorm"
)

func (rs *RdbService) CreateGroup(groups []*model.Group) error {
	return rs.tx.Create(&groups).Error
}

func (rs *RdbService) IncrGroupSeq(groupID int64) (err error) {
	return rs.tx.Table("groups").Where("group_id = ?", groupID).Update("max_seq", gorm.Expr("max_seq + 1")).Error
}

func (rs *RdbService) GetGroupMember(groupID int64) (members []model.GroupUser, err error) {
	err = rs.tx.Table("groups").Where("group_id = ?", groupID).Find(&members).Error
	return
}
