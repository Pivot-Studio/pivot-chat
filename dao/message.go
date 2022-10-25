package dao

import "github.com/Pivot-Studio/pivot-chat/model"


func (rs *RdbService) CreateMessage(meg []*model.Message) error {
	return rs.tx.Create(&meg).Error
}