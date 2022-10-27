package api

import (
	"github.com/sirupsen/logrus"
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
type chgPwdParam struct {
	OldPwd   string `form:"oldpwd" binding:"required"`
	NewPwd   string `form:"newpwd" binding:"required"`
	UserName string `form:"username" binding:"required"`
}

type emailParam struct {
	Email string `form:"email" binding:"required"`
}

func ChgPwd(ctx *gin.Context) {
	p := &chgPwdParam{}
	err := ctx.ShouldBind(p)
	if err != nil {
		logrus.Errorf("[chgPwd] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "修改密码失败",
		})
		return
	}
	oldPwdHash, err := util.EncodePassword(p.OldPwd)
	if err != nil {
		logrus.Errorf("[chgPwd] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "修改密码失败",
		})
		return
	}
	newPwdHash, err := util.EncodePassword(p.NewPwd)
	if err != nil {
		logrus.Errorf("[chgPwd] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "修改密码失败",
		})
		return
	}

	err = service.ChgPwd(ctx, p.UserName, oldPwdHash, newPwdHash)

	if err != nil {
		logrus.Errorf("[chgPwd] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "修改密码失败",
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"meg": "修改密码成功",
	})
}

func Register(ctx *gin.Context) {
	p := &registerParam{}
	err := ctx.ShouldBind(p)
	if err != nil {
		logrus.Errorf("[Register] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "注册失败",
		})
		return
	}

	passwordHash, err := util.EncodePassword(p.Password)
	if err != nil {
		logrus.Errorf("[Register] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "注册失败",
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
		logrus.Errorf("[Register] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "注册失败",
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

func Email(ctx *gin.Context) {
	p := &emailParam{}
	err := ctx.ShouldBind(p)
	if err != nil {
		logrus.Errorf("[Email] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "邮箱或密码格式不合法",
		})
		return
	}
	code := service.CreatCode()
	err = service.SendEmail(ctx, p.Email, code)
	if err != nil {
		logrus.Errorf("[Email] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "验证码发送失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送验证码成功",
	})
}
