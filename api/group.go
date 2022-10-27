package api

import (
	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/Pivot-Studio/pivot-chat/service"
	"github.com/sirupsen/logrus"
)

func HandleGroupMessage(meg *model.GroupMessageInput) {
	if !dao.RS.ExistGroup(meg.GroupId) {
		return
	}
	// group := service.GetUpdatedGroup(meg.GroupId)
	err := service.SendMessage(meg)
	if err != nil {
		logrus.Fatalf("[HandleGroupMessage] SendMessage %+v", err)
	}
}
