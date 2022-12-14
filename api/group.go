package api

import (
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/Pivot-Studio/pivot-chat/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type getMembersByGroupIdParam struct {
	GroupId int64 `form:"groupid" binding:"required"`
}

func HandleGroupMessage(meg *model.GroupMessageInput) error {
	err := service.GroupOp.SaveGroupMessage(meg)
	return err
}

func HandleJoinGroup(meg *model.UserJoinGroupInput) error {
	err := service.GroupOp.JoinGroup(meg)
	return err
}

func GetMembersByGroupId(ctx *gin.Context) {

	p := &getMembersByGroupIdParam{}
	err := ctx.ShouldBind(p)
	if err != nil {
		logrus.Errorf("[GetMembersByGroupId] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "参数解析错误，查询失败",
		})
		return
	}

	data, err := service.GroupOp.GetMembersByGroupId(ctx, p.GroupId)

	if err != nil {
		logrus.Errorf("[GetMembersByGroupId] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error() + "，查询失败",
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"data": data,
	})

}

type CreateGroupParams struct {
	Name         string `json:"group_name"`
	Introduction string `json:"introduction"`
}

func CreateGroup(ctx *gin.Context) {
	p := &CreateGroupParams{}
	err := ctx.ShouldBindJSON(p)
	resp := &service.CreateGroupResp{}
	if err != nil || p.Name == "" {
		logrus.Errorf("[api.CreateGroup] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg":  "创建失败, 参数不合法",
			"data": *resp,
		})
	}

	user, err := service.GetUserFromAuth(ctx)
	if err != nil {
		logrus.Errorf("[api.CreateGroup] GetUserFromAuth %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg":  "创建失败, 鉴权失败",
			"data": *resp,
		})
	}

	resp, err = service.CreateGroup(p.Name, p.Introduction, user.UserId)
	if err != nil {
		logrus.Errorf("[api.CreateGroup] %+v", err)
		resp = &service.CreateGroupResp{}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg":  "创建失败, 服务器错误",
			"data": *resp,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "创建成功",
		"data": *resp,
	})
}
