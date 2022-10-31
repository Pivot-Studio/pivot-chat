package dao

import "github.com/Pivot-Studio/pivot-chat/model"

func (rs *RdbService) CreateMessage(meg []*model.Message) error {
	return rs.tx.Create(&meg).Error
}

func (rs *RdbService) SyncMessage(receiverId int64, syncSeq int64, limit int, isNew bool) (megs []model.Message, err error) {
	if isNew {
		err = rs.tx.Table("messages").Where("receiver_id = ? AND Seq <= ?", receiverId, syncSeq).Order("seq DESC").Limit(limit).Find(&megs).Error
	} else {
		err = rs.tx.Table("messages").Where("receiver_id = ? AND Seq >= ?", receiverId, syncSeq).Order("seq ASC").Limit(limit).Find(&megs).Error
	}
	return megs, err
}
