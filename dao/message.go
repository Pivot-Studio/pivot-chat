package dao

import (
	"github.com/Pivot-Studio/pivot-chat/model"
)

func (rs *RdbService) CreateMessage(meg []*model.Message) error {
	return rs.tx.Create(&meg).Error
}

func (rs *RdbService) SyncMessage(receiverId int64, syncSeq int64, limit int) (megs []model.Message, err error) {
	err = rs.tx.Table("messages").Where("receiver_id = ? AND Seq >= ?", receiverId, syncSeq).Order("seq").Limit(limit).Find(&megs).Error
	return megs, err
}
