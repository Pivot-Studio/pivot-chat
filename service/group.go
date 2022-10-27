package service

import (
	"encoding/json"
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
	SenderId    int64  // 用户id
	SenderType  int64  // 发送者身份
	Message     string // 消息内容
	ReceiverId  int64  // 群组id
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

	if !IsMember(sendInfo.UserId, members) {
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
	for _, user := range members {
		// 前面已经发送过，这里不需要再发送
		// if sendInfo.SenderType == SenderType_USER && user.UserId == sendInfo.SenderId {
		// 	continue
		// }
		user0 := user
		go func() {
			output := model.GroupMessageOutput{ 
				UserId:   user0.UserId,
				GroupId:  g.GroupId,
				Data:     sendInfo.Data,
				SenderId: sendInfo.UserId,
				Seq:      g.MaxSeq + 1,
			}
			bytes, err := json.Marshal(output)
			if err != nil {
				logrus.Fatalf("[Service] | conn-manager json Marshal err:", err)
				return
			}
			err = SendToUser(user0.UserId, bytes, PackageType_PT_MESSAGE)
			if err != nil {
				logrus.Fatalf("[Service] | group sendmeg error:", err)
				return
			}
		}()
	}
	return nil
}

// func HandleGroupMessage(meg *model.Message) {
// 	if !dao.RS.ExistGroup(meg.ReceiverId) {
// 		return
// 	}
// 	group := GetUpdatedGroup(meg.ReceiverId)
// 	err := group.SendMessage(SendInfo{
// 		SenderId:     meg.SenderId,
// 		Message:    string(meg.Content),
// 		SenderType: meg.SenderType,
// 	})
// 	if err != nil {
// 		logrus.Fatalf("[HandleGroupMessage] SendMessage %+v", err)
// 	}
// }
