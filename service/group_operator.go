package service

import (
	"encoding/json"
	"github.com/Pivot-Studio/pivot-chat/constant"
	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var GroupOp GroupOperator

type Group_ struct {
	group   *model.Group
	Members *[]model.GroupUser
	sync.Mutex
}
type GroupOperator struct {
	Groups    *[]Group_
	GroupsMap sync.Map
	lock      sync.Mutex
}

func (g *Group_) IsMember(userID int64) bool {
	//todo
	for i := range *g.Members {
		if (*g.Members)[i].UserId == userID {
			return true
		}
	}
	return false
}
func (gpo *GroupOperator) StoreGroup(groupID int64, group *Group_) {
	GroupOp.GroupsMap.Store(groupID, group)
}

// GetGroup 根据groupID获取群组, 能够保证获取的群组成员是最新的, 其他属性更新时需要保持DB中一致性
func (gpo *GroupOperator) GetGroup(groupID int64) (*Group_, error) {
	//全局锁
	gpo.lock.Lock()
	value, ok := gpo.GroupsMap.Load(groupID)
	if !ok {
		//从数据库中查找
		g, err := dao.RS.QueryGroup(groupID)
		if err != nil {
			gpo.lock.Unlock()
			logrus.Errorf("[service.GetGroup] QueryGroup %+v", err)
			return nil, constant.NotGroupRecordErr
		}
		//缓存Group
		value = &Group_{
			group: g,
		}
		gpo.StoreGroup(groupID, value.(*Group_))
	}
	gpo.lock.Unlock()

	//对群组所有写操作, 必须锁住对应的群组
	g := value.(*Group_)
	//更新群组成员
	g.Lock()
	if g.group.UserNum != int32(len(*g.Members)) {
		var err error
		g.Members, err = dao.RS.GetGroupUsers(g.group.GroupId)
		if err != nil {
			logrus.Errorf("[service.SaveGroupMessage] GetGrupUsers %+v", err)
			g.Unlock()
			return g, constant.GroupGetMembersErr
		}
	}
	g.Unlock()

	return g, nil
}
func (g *Group_) SendGroupMessage(sendInfo *model.GroupMessageInput, seq int64) {
	// 将消息发送给群组用户
	// 复制一份以免遍历时改变group导致错误
	members := *g.Members
	for _, user := range members {
		userID := user.UserId
		go func() {
			output := model.GroupMessageOutput{
				UserId:   userID,
				GroupId:  g.group.GroupId,
				Data:     sendInfo.Data,
				SenderId: sendInfo.UserId,
				Seq:      seq,
			}

			bytes, err := json.Marshal(output)
			if err != nil {
				logrus.Fatalf("[service.SendGroupMessage] json Marshal %+v", err)
				return
			}

			err = SendToUser(userID, bytes, PackageType_PT_MESSAGE)
			if err != nil {
				logrus.Fatalf("[service.SendGroupMessage] group SendToUser %+v", err)
				return
			}
		}()
	}
}

// JoinGroup todo
func (gpo *GroupOperator) JoinGroup(groupID int64, user *model.GroupUser) error {
	g, err := gpo.GetGroup(groupID)
	if err != nil {
		logrus.Errorf("[service.JoinGroup] GetGroup %+v", err)
		return err
	}

	g.Lock()
	if g.IsMember(user.UserId) {
		return nil
	}
	//todo
	*g.Members = append(*g.Members, *user)
	g.group.UserNum += 1

	g.Unlock()
	return nil
}

// SaveGroupMessage 持久化群组消息
func (gpo *GroupOperator) SaveGroupMessage(SendInfo *model.GroupMessageInput, g *Group_) error {
	g, err := gpo.GetGroup(SendInfo.GroupId)
	if err != nil {
		logrus.Errorf("[service.SaveGroupMessage] GetGroup %+v", err)
		return constant.NotGroupRecordErr
	}
	/*
		这里不加锁:理由有二:
				 对接受方而言, 加入群组时, 会收到并发消息, 因为g是共享的
				 			 退出群组时, 不会收到消息(如果在Send时退出)
		   		 对发送方而言, 加入群组时, 并发的发消息, 并不存在这种情况
							 退出群组时, 发不出消息(如果在Send时退出)
	*/
	if !g.IsMember(SendInfo.UserId) {
		logrus.Errorf("[service.SaveGroupMessage] IsMember")
		return constant.UserNotMatchGroup
	}

	//开始持久化
	meg := &model.Message{
		SenderId:   SendInfo.UserId,
		ReceiverId: SendInfo.GroupId,
		Content:    SendInfo.Data,
		SendTime:   time.Now(),
	}
	//保证MaxSeq是正确的, 需要加锁
	g.Lock()
	meg.Seq = g.group.MaxSeq + 1

	err = dao.RS.IncrGroupSeq(g.group.GroupId)
	if err != nil {
		logrus.Errorf("[Service.SaveGroupMessage] IncrGroupSeq %+v", err)
		return err
	}
	//持久化到DB
	err = dao.RS.CreateMessage([]*model.Message{meg})
	if err != nil {
		logrus.Errorf("[Service.SaveGroupMessage] CreateMessage %+v", err)
		g.Unlock()
		return err
	}
	g.Unlock()

	//发送消息, 发送的成员是按现在Group的成员(可能被改变)
	g.SendGroupMessage(SendInfo, meg.Seq)
	return nil
}
