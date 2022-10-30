package api

import (
	"errors"

	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/Pivot-Studio/pivot-chat/service"
)

func HandleGroupMessage(meg *model.GroupMessageInput) error {
	if !dao.RS.ExistGroup(meg.GroupId) {
		return errors.New("group not existed!")
	}
	err := service.GroupOp.SaveGroupMessage(meg)
	return err
}

func HandleJoinGroup(meg *model.UserJoinGroupInput) error {
	if !dao.RS.ExistGroup(meg.GroupId) {
		return errors.New("group not existed!")
	}
	err := service.GroupOp.JoinGroup(meg)
	return err
}
