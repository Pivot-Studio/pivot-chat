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

func (rs *RdbService) ExistGroup(groupID int64) bool {
	return rs.tx.Table("groups").Where("group_id = ?", groupID).
		Take(model.Group{}).Error != gorm.ErrRecordNotFound
}

func (rs *RdbService) QueryGroup(groupID int64) (*model.Group, error) {
	g := model.Group{}
	err := rs.tx.Table("groups").Where("group_id = ?", groupID).Take(&g).Error
	return &g, err
}

func (rs *RdbService) GetGroupUsers(groupID int64) (*[]model.GroupUser, error) {
	var g []model.GroupUser
	err := rs.tx.Table("groups").Where("group_id = ?", groupID).Find(&g).Error
	return &g, err
}
