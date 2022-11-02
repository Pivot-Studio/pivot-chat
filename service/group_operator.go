package service

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/Pivot-Studio/pivot-chat/constant"
	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var GroupOp GroupOperator

type Group_ struct {
	group   *model.Group
	Members *[]model.GroupUser
	sync.RWMutex
}
type GroupOperator struct {
	Groups    []*Group_
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

// SendGroupMessage 群发消息, 无锁
func (g *Group_) SendGroupMessage(sendInfo *model.GroupMessageInput, seq int64) {
	// 将消息发送给群组用户
	// 复制一份以免遍历时改变group导致错误, 这里也可以考虑加锁, 但是这样会更快一点
	var members []model.GroupUser
	copy(members, *g.Members)
	for _, user := range members {
		user0 := user
		go func(user *model.GroupUser, sendInfo *model.GroupMessageInput) {
			output := model.GroupMessageOutput{
				UserId:   user0.UserId,
				GroupId:  g.group.GroupId,
				Data:     sendInfo.Data,
				SenderId: sendInfo.UserId,
				Seq:      seq,
				ReplyTo:  sendInfo.ReplyTo,
				Type:     sendInfo.Type,
			}

			bytes, err := json.Marshal(output)
			if err != nil {
				logrus.Fatalf("[service.SendGroupMessage] json Marshal %+v", err)
				return
			}

			err = SendToUser(user0.UserId, bytes, PackageType_PT_MESSAGE)
			if err != nil {
				logrus.Fatalf("[service.SendGroupMessage] group SendToUser %+v", err)
				return
			}
		}(&user0, sendInfo)
	}
}

// UpdateGroup todo 修改群组信息
func (gpo *GroupOperator) UpdateGroup() {

}

// QuitGroup todo 退出群组
func (gpo *GroupOperator) QuitGroup() {

}

// JoinGroup 加入群组
func (gpo *GroupOperator) JoinGroup(input *model.UserJoinGroupInput) error {
	g, err := gpo.GetGroup(input.GroupId)
	if err != nil {
		logrus.Errorf("[service.JoinGroup] GetGroup %+v", err)
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

	g.Lock()
	defer g.Unlock()
	if g.IsMember(input.UserId) {
		return nil
	}
	err = dao.RS.CreateGroupUser([]*model.GroupUser{&groupUser})
	if err != nil {
		return err
	}
	err = dao.RS.IncrGroupUserNum(g.group.GroupId)
	if err != nil {
		return err
	}
	// 缓存放最后更新, 保证缓存与数据库同步
	*g.Members = append(*g.Members, groupUser)
	g.group.UserNum += 1

	// 给加入的用户回复消息
	output := model.UserJoinGroupOutput{
		GroupId:      g.group.GroupId,
		OwnerId:      g.group.OwnerId,
		Name:         g.group.Name,
		Introduction: g.group.Introduction,
		UserNum:      g.group.UserNum,
		CreateTime:   g.group.CreateTime,
	}
	err = SendToUser(input.UserId, output, PackageType_PT_JOINGROUP)
	if err != nil {
		logrus.Fatalf("[Service] UserJoinGroup %+v", err)
		return err
	}
	return nil
}

// SaveGroupMessage 持久化群组消息, 同时会发送给每一个人
func (gpo *GroupOperator) SaveGroupMessage(SendInfo *model.GroupMessageInput) error {
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

// GetMembersByGroupId GetUsersByGroupId 根据GroupID获取当前members
func (gpo *GroupOperator) GetMembersByGroupId(ctx *gin.Context, groupID int64) ([]map[string]interface{}, error) {
	_, err := GetUserFromAuth(ctx)
	if err != nil {
		return nil, err
	}
	g, err := gpo.GetGroup(groupID)
	if err != nil {
		logrus.Errorf("[service.GetMembersByGroupId] GetGroup %+v", err)
		return nil, constant.GroupGetMembersErr
	}
	// copy一遍以免遍历出现并发问题
	var members []model.GroupUser
	copy(members, *g.Members)

	ret := make([]map[string]interface{}, 0)
	for _, member := range members {
		data := make(map[string]interface{})
		user := &model.User{}
		user.UserId = member.UserId
		err = dao.RS.GetUserbyId(user)
		if err != nil {
			logrus.Errorf("[service.GetMembersByGroupId] GetUserbyId %+v", err)
			continue
		}
		data["user_name"] = user.UserName
		data["user_id"] = user.UserId
		data["email"] = user.Email
		data["type_in_group"] = member.MemberType
		if err != nil {
			logrus.Errorf("[service.GetMembersByGroupId] GetGroup %+v", err)
			return nil, constant.GroupGetMembersErr
		}
		ret = append(ret, data)
	}
	return ret, nil
}

type CreateGroupResp struct {
	GroupId      int64     `json:"group_id"`
	OwnerId      int64     `json:"owner_id"`
	Name         string    `json:"name"`
	Introduction string    `json:"introduction"`
	UserNum      int32     `json:"user_num"`
	CreateTime   time.Time `json:"create_time"`
	MaxSeq       int64     `json:"max_seq"`
}

func CreateGroup(Name string, Introduction string, OwnerId int64) (*CreateGroupResp, error) {
	g := &model.Group{
		OwnerId:      OwnerId,
		Name:         Name,
		Introduction: Introduction,
		UserNum:      1,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		MaxSeq:       0,
	}

	err := dao.RS.CreateGroup(g)
	if err != nil {
		logrus.Errorf("[service] CreateGroup %+v", err)
		return nil, err
	}

	resp := &CreateGroupResp{
		GroupId:      g.GroupId,
		OwnerId:      g.OwnerId,
		Name:         g.Name,
		Introduction: g.Introduction,
		UserNum:      g.UserNum,
		CreateTime:   g.CreateTime,
		MaxSeq:       g.MaxSeq,
	}
	return resp, nil
}
