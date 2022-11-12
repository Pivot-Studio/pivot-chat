package api

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/Pivot-Studio/pivot-chat/service"
	"github.com/Pivot-Studio/pivot-chat/util"
	"github.com/gin-gonic/gin"
)

type registerParam struct {
	UserName string `form:"user_name" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Captcha  string `form:"captcha" binding:"required"`
}
type chgPwdParam struct {
	OldPwd string `form:"oldpwd" binding:"required"`
	NewPwd string `form:"newpwd" binding:"required"`
	Email  string `form:"email" binding:"required"`
}

type emailParam struct {
	Email string `form:"email" binding:"required"`
}

type findUserByIdParam struct {
	UserId int64 `form:"user_id" binding:"required"`
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
	passwordHash, err := util.EncodePassword(p.NewPwd)
	if err != nil {
		logrus.Errorf("[Register] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "密码hash失败，注册失败",
		})
		return
	}
	err = service.ChgPwd(p.Email, p.OldPwd, passwordHash)

	if err != nil {
		logrus.Errorf("[chgPwd] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "修改密码失败:" + err.Error(),
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "修改密码成功",
	})
}

func FindUserById(ctx *gin.Context) {
	p := &findUserByIdParam{}
	err := ctx.ShouldBind(p)
	if err != nil {
		logrus.Errorf("[FindUserById] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "参数解析错误，查询失败",
		})
		return
	}

	data, err := service.FindUserById(ctx, p.UserId)

	if err != nil {
		logrus.Errorf("[FindUserById] %+v", err)
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
	err = service.Register(&model.User{
		Password: passwordHash,
		Email:    p.Email,
		UserName: p.UserName,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}, p.Captcha)
	if err != nil {
		logrus.Errorf("[Register] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "注册失败，" + err.Error(),
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
	code, err := service.Email(ctx, p.Email)
	// err = dao.Cache.Set(context.Background(), p.Email, code, time.Minute*30).Err()
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

type loginParam struct {
	Password string `form:"password" binding:"required"`
	Email    string `form:"email" binding:"required"`
}

func Login(ctx *gin.Context) {
	p := &loginParam{}
	err := ctx.ShouldBind(p)
	if err != nil {
		logrus.Errorf("[Login] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "邮箱或密码格式不合法",
		})
		return
	}

	user, token, err := service.Login(p.Email, p.Password)
	if err != nil {
		logrus.Errorf("[Login] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "登录失败" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
		"data": gin.H{
			"token":     token,
			"user_id":   user.UserId,
			"user_name": user.UserName,
			"email":     user.Email,
		},
	})

}

func GetMyGroups(ctx *gin.Context) {
	user, err := service.GetUserFromAuth(ctx)
	if err != nil {
		logrus.Errorf("[api.GetMyGroups] GetUserFromAuth %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "查询失败，" + err.Error(),
		})
		return
	}

	g, err := service.GetMyGroups(user.UserId)

	if err != nil {
		logrus.Errorf("[api.GetMyGroups] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg":  "查询失败",
			"data": *g,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"data": *g,
	})

}

func GetMyJoinedGroups(ctx *gin.Context) {
	user, err := service.GetUserFromAuth(ctx)
	if err != nil {
		logrus.Errorf("[api.GetMyGroups] GetUserFromAuth %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "查询失败，" + err.Error(),
		})
		return
	}

	g, err := service.GetMyJoinedGroups(user.UserId)

	if err != nil {
		logrus.Errorf("[api.GetMyGroups] %+v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg":  "查询失败",
			"data": *g,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"data": *g,
	})

}
