package api

import (
	"net/http"
	"time"

	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/Pivot-Studio/pivot-chat/service"
	"github.com/Pivot-Studio/pivot-chat/util"
	"github.com/gin-gonic/gin"
)

type registerParam struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Captcha  string `form:"captcha" binding:"required"`
}

func Register(ctx *gin.Context) {
	p := &registerParam{}
	err := ctx.ShouldBind(p)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "[Register]:" + err.Error(),
		})
		return
	}
	passwordHash, err := util.EncodePassword(p.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "[Register]:" + err.Error(),
		})
		return
	}
	err = service.Register(ctx, &model.User{
		Password: passwordHash,
		Email:    p.Email,
		UserName: p.UserName,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}, p.Captcha)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "[Register]:" + err.Error(),
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}
