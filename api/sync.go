package api

import (
	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/Pivot-Studio/pivot-chat/service"
	"github.com/sirupsen/logrus"
)

func HandleSync(meg *model.GroupMessageSyncInput) {
	if !dao.RS.ExistGroup(meg.GroupId) {
		return
	}
	err := service.Sync(meg)
	if err != nil {
		logrus.Fatalf("[HandleSync] service.Sync error %+v", err)
	}
}