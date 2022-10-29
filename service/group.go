package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/Pivot-Studio/pivot-chat/constant"
	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/sirupsen/logrus"
)

// Group 群组
type Group struct {
	*model.Group
}

var lock sync.Mutex

type SendInfo struct {
	SenderId   int64  // 用户id
	SenderType int64  // 发送者身份
	Message    string // 消息内容
	ReceiverId int64  // 群组id
}

const (
	SenderType_USER    = 1
	ReceiverType_USER  = 2
	ReceiverType_GROUP = 3
)

func IsMember(userId int64, members []model.GroupUser) bool {
	for i := range members {
		if members[i].UserId == userId {
			return true
		}
	}
	return false
}

func SendMessage(sendInfo *model.GroupMessageInput) error { // 进入这里时，group内容是跟数据库一致的，members也是使用的正确缓存
	lock.Lock()
	defer lock.Unlock()
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered. Error:\n", r)
	// 	}
	// }()
	g, err := GetUpdatedGroup(sendInfo.GroupId) // 这肯定是最新的，而且是一次
	if err != nil {
		return err
	}
	members, err := dao.RS.GetGroupUsers(g.GroupId)
	if err != nil {
		return err
	}

	if !IsMember(sendInfo.UserId, *members) {
		logrus.Fatalf("[Service] | group sendmeg error: user isn't in group | sendInfo:", sendInfo)
		return constant.UserNotMatchGroup
	}
	// 持久化
	meg := model.Message{
		SenderId:   sendInfo.UserId,
		ReceiverId: sendInfo.GroupId,
		Content:    sendInfo.Data,
		Seq:        g.MaxSeq + 1,
		SendTime:   time.Now(),
		Type:       sendInfo.Type,
		ReplyTo:    sendInfo.ReplyTo,
	}
	err = dao.RS.CreateMessage([]*model.Message{&meg})
	if err != nil {
		logrus.Fatalf("[Service] | conn-manager persist CreateMessage err:", err)
		return err
	}
	// 持久化消息成功 update group
	err = dao.RS.IncrGroupSeq(g.GroupId)
	if err != nil {
		logrus.Fatalf("[Service] | conn-manager persist IncrGroupSeq err:", err)
		return err
	}

	// 将消息发送给群组用户
	for _, user := range *members {
		// 前面已经发送过，这里不需要再发送
		// if sendInfo.SenderType == SenderType_USER && user.UserId == sendInfo.SenderId {
		// 	continue
		// }
		user0 := user
		go func(user *model.GroupUser, sendInfo *model.GroupMessageInput) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered. Error:\n", r)
				}
			}()
			output := model.GroupMessageOutput{
				UserId:   user.UserId,
				GroupId:  g.GroupId,
				Data:     sendInfo.Data,
				SenderId: sendInfo.UserId,
				Seq:      g.MaxSeq + 1,
				ReplyTo:  sendInfo.ReplyTo,
				Type:     sendInfo.Type,
			}
			err = SendToUser(user.UserId, output, PackageType_PT_MESSAGE)
			if err != nil {
				logrus.Fatalf("[Service] | group sendmeg error:", err)
				return
			}
		}(&user0, sendInfo)
	}
	return nil
}

func UserJoinGroup(input *model.UserJoinGroupInput) error {
	lock.Lock()
	defer lock.Unlock()
	g, err := GetUpdatedGroup(input.GroupId) // 这肯定是最新的，而且是一次
	if err != nil {
		return err
	}
	groupUser := model.GroupUser{
		GroupId:    input.GroupId,
		UserId:     input.UserId,
		MemberType: model.SPEAKER,
		Status:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err = dao.RS.CreateGroupUser([]*model.GroupUser{&groupUser})
	if err != nil {
		return err
	}
	err = dao.RS.IncrGroupUserNum(g.GroupId)
	if err != nil {
		return err
	}
	output := model.UserJoinGroupOutput{
		UserId:       input.UserId,
		GroupId:      g.GroupId,
		Name:         g.Name,
		Introduction: g.Introduction,
		UserNum:      g.UserNum,
		CreateTime:   g.CreateTime,
	}
	err = SendToUser(input.UserId, output, PackageType_PT_JOINGROUP)
	if err != nil {
		logrus.Fatalf("[Service] | UserJoinGroup error:", err)
		return err
	}
	return nil
}
