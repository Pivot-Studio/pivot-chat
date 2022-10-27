package service

import (
	"encoding/json"

	"github.com/Pivot-Studio/pivot-chat/constant"
	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/sirupsen/logrus"
)

func Sync(input *model.GroupMessageSyncInput) error { // 进入这里时，group内容是跟数据库一致的，members也是使用的正确缓存
	lock.Lock()
	defer lock.Unlock()
	g, err := GetUpdatedGroup(input.GroupId) // 这肯定是最新的，而且是一次
	if err != nil {
		return err
	}
	members, err := dao.RS.GetGroupUsers(g.GroupId)
	if err != nil {
		return err
	}
	g, err = GetUpdatedGroup(input.GroupId)
	if !IsMember(input.UserId, members) {
		logrus.Fatalf("[Service] | sync error: user isn't in group | input:", input)
		return constant.UserNotMatchGroup
	}
	megs, err := dao.RS.SyncMessage(input.GroupId, input.SyncSeq)
	if err != nil {
		return err
	}
	groupMessageOutput := make([]model.GroupMessageOutput, 0)
	for _, meg := range megs {
		groupMessageOutput = append(groupMessageOutput, model.GroupMessageOutput{
			UserId:   input.UserId,
			GroupId:  meg.ReceiverId,
			Data:     meg.Content,
			SenderId: meg.SenderId,
			Seq:      meg.Seq,
		})
	}
	output := model.GroupMessageSyncOutput{
		UserId:  input.UserId,
		GroupId: input.GroupId,
		Data:    groupMessageOutput,
		MaxSeq:  megs[len(megs)-1].Seq,
	}
	bytes, err := json.Marshal(output)
	if err != nil {
		return err
	}
	err = SendToUser(input.UserId, bytes, PackageType_PT_SYNC)
	if err != nil {
		return err
	}
	return nil
}
