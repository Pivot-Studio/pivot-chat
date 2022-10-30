package api

import (
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/Pivot-Studio/pivot-chat/service"
)

func HandleGroupMessage(meg *model.GroupMessageInput) error {
	err := service.GroupOp.SaveGroupMessage(meg)
	return err
}

func HandleJoinGroup(meg *model.UserJoinGroupInput) error {
	err := service.GroupOp.JoinGroup(meg)
	return err
}
