package api

import (
	"net/http"

	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/Pivot-Studio/pivot-chat/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
