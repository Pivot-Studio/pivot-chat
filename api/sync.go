package api

import (
	"fmt"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/Pivot-Studio/pivot-chat/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SyncParam struct {
	GroupId int64 `form:"group_id" binding:"required"` // 群组id
	SyncSeq int64 `form:"sync_seq" binding:"required"` // 开始同步的seq，是用户的本地seq+1
	Limit   int64 `form:"limit" binding:"required"`
	IsNew   int64 `form:"is_new" binding:"required"`
}

func Sync(ctx *gin.Context) {
	p := &SyncParam{}
	err := ctx.ShouldBind(p)
	if err != nil {
		logrus.Errorf("[Sync] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "sync参数绑定错误",
		})
		return
	}
	input := &model.GroupMessageSyncInput{
		GroupId: p.GroupId,
		SyncSeq: p.SyncSeq,
		Limit:   p.Limit,
		IsNew:   p.IsNew,
	}
	ret, err := service.Sync(ctx, input)
	if err != nil {
		if err != nil {
			logrus.Errorf("[Sync] %+v", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "Sync err:" + err.Error(),
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("同步%d条消息成功", len(ret.Data)),
		"data": ret,
	})
	return
}
