package service

import (

	"github.com/Pivot-Studio/pivot-chat/constant"
	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/Pivot-Studio/pivot-chat/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Sync(ctx *gin.Context, input *model.GroupMessageSyncInput) (*model.GroupMessageSyncOutput, error) { // 进入这里时，group内容是跟数据库一致的，members也是使用的正确缓存
	user, err := GetUserFromAuth(ctx)
	if err != nil {
		return nil, err
	}
	input.UserId = user.UserId
	g, err := GroupOp.GetGroup(input.GroupId)
	if err != nil {
		return nil, err
	}
	if !g.IsMember(input.UserId) {
		logrus.Errorf("[Service] | sync error: user isn't in group | input:%+v", input)
		return nil, constant.UserNotMatchGroup
	}
	megs, err := dao.RS.SyncMessage(input.GroupId, input.SyncSeq, int(input.Limit), input.IsNew)
	if err != nil {
		return nil, err
	}
	// if len(megs) <= 0 {
	// 	return nil, errors.New("未找到符合条件的消息")
	// }
	groupMessageOutput := make([]model.GroupMessageOutput, 0)
	for _, meg := range megs {
		content, err := util.AESdecrypt(meg.Content)
		if err != nil {
			logrus.Errorf("[Service.Sync] 消息解密失败" + err.Error())
			continue
		}
		groupMessageOutput = append(groupMessageOutput, model.GroupMessageOutput{
			UserId:   input.UserId,
			GroupId:  meg.ReceiverId,
			Data:     content,
			SenderId: meg.SenderId,
			Seq:      meg.Seq,
			ReplyTo:  meg.ReplyTo,
			Type:     meg.Type,
			Time:     meg.SendTime,
		})
	}
	if len(megs) <= 0 {
		output := model.GroupMessageSyncOutput{
			UserId:  input.UserId,
			GroupId: input.GroupId,
			Data:    groupMessageOutput,
			MaxSeq:  0,
		}
		return &output, nil
	}
	output := model.GroupMessageSyncOutput{
		UserId:  input.UserId,
		GroupId: input.GroupId,
		Data:    groupMessageOutput,
		MaxSeq:  megs[len(megs)-1].Seq,
	}

	return &output, nil
}
