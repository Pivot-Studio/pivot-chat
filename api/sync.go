package api

import (
	"errors"

	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/Pivot-Studio/pivot-chat/service"
	"github.com/sirupsen/logrus"
)

func HandleSync(meg *model.GroupMessageSyncInput) error {
	if !dao.RS.ExistGroup(meg.GroupId) {
		return errors.New("group not existed!")
	}
	err := service.Sync(meg)
	if err != nil {
		logrus.Fatalf("[HandleSync] service.Sync error %+v", err)
		return err
	}
	return nil
}