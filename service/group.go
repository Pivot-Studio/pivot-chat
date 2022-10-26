package service

import (
	"encoding/json"
	"fmt"
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

type SendInfo struct {
	ReceiverId int64  // 用户id
	Messgae    string // 消息内容
	SenderType int64  // 发送者身份，即群聊id
	SenderId   int64  // 发送者id
}

const (
	SenderType_USER    = 1
	ReceiverType_USER  = 2
	ReceiverType_GROUP = 3
)

func (g *Group) GetGroupMember() ([]*model.GroupUser, error) {
	members, err := dao.RS.GetGroupMember(g.GroupId)
	if err != nil {
		return nil, err
	}
	membersPointers := make([]*model.GroupUser, 0)
	for _, member := range members {
		membersPointers = append(membersPointers, &member)
	}
	return membersPointers, nil
}

func (g *Group) IsMember(userId int64) bool {
	for _, member := range g.Members {
		if member.GroupUserId == userId {
			return true
		}
	}
	return false
}

func (g *Group) SendMessgae(sendInfo SendInfo) error {
	if sendInfo.SenderType == SenderType_USER && !g.IsMember(sendInfo.ReceiverId) {
		logrus.Fatalf("[Service] | group sendmeg error: user isn't in group | sendInfo:", sendInfo)
		return constant.UserNotMatchGroup
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered. Error:\n", r)
			}
		}()
		bytes, err := json.Marshal(sendInfo.Messgae)
		if err != nil {
			logrus.Fatalf("[Service] | conn-manager json Marshal err:", err)
			return
		}
		// 持久化
		meg := model.Message{
			SenderType:   sendInfo.SenderType,
			SenderId:     sendInfo.SenderId,
			ReceiverType: ReceiverType_GROUP,
			ReceiverId:   g.GroupId,
			Content:      bytes,
			Seq:          g.MaxSeq + 1,
			SendTime:     time.Now(),
		}
		err = dao.RS.CreateMessage([]*model.Message{&meg})
		if err != nil {
			logrus.Fatalf("[Service] | conn-manager persist CreateMessage err:", err)
			return
		}
		// 持久化消息成功 update group
		err = dao.RS.IncrGroupSeq(g.GroupId)
		if err != nil {
			logrus.Fatalf("[Service] | conn-manager persist IncrGroupSeq err:", err)
			return
		}
		// 将消息发送给群组用户
		for _, user := range g.Members {
			// 前面已经发送过，这里不需要再发送
			if sendInfo.SenderType == SenderType_USER && user.UserId == sendInfo.SenderId {
				continue
			}

			err = SendToUser(sendInfo.ReceiverId, bytes)
			if err != nil {
				logrus.Fatalf("[Service] | group sendmeg error:", err)
				continue
			}
		}
	}()
	return nil
}
