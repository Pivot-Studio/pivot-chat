package dao

import (
	"fmt"

	"github.com/Pivot-Studio/pivot-chat/model"
)

func (rs *RdbService) CreateMessage(meg []*model.Message) error {
	return rs.tx.Create(&meg).Error
}

func (rs *RdbService) SyncMessage(receiverId int64, syncSeq int64, limit int, isNew int64) (megs []model.Message, err error) {
	fmt.Println("ISNEW: ", isNew)
	if isNew > 0 {
		tmp := rs.tx.Debug().Table("messages").Where("receiver_id = ? AND Seq >= ?", receiverId, syncSeq).Order("seq desc").Limit(limit)
		err = tmp.Order("seq").Find(&megs).Error
	} else {
		err = rs.tx.Table("messages").Where("receiver_id = ? AND Seq >= ?", receiverId, syncSeq).Order("seq").Limit(limit).Find(&megs).Error
	}
	return megs, err
}
