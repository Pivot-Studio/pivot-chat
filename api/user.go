package api

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Pivot-Studio/pivot-chat/dao"
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
			"msg": "参数绑定错误，修改密码失败",
		})
		return
	}

	err = service.ChgPwd(ctx, p.UserName, p.NewPwd)

	if err != nil {
		logrus.Errorf("[chgPwd] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "修改密码失败:" + err.Error(),
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
			"msg": "参数解析错误，注册失败",
		})
		return
	}

	passwordHash, err := util.EncodePassword(p.Password)
	if err != nil {
		logrus.Errorf("[Register] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "密码hash失败，注册失败",
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
			"msg": err.Error(),
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
	err = dao.Cache.Set(context.Background(), p.Email, code, time.Minute*30).Err()
	go func() {
		emailctx, canal := context.WithTimeout(context.TODO(), 3*time.Second)
		defer canal()
		err = service.SendEmail(emailctx, p.Email, code)
	}()
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
