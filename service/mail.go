package service

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Pivot-Studio/pivot-chat/conf"
	"github.com/Pivot-Studio/pivot-chat/dao"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

var (
	d            *gomail.Dialer
	emailContent string = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8"/>
		<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
		<title>Document</title>
	</head>
	<style>
		* {
			margin: 0;
			padding: 0;
		}
	
		.main {
			margin: auto;
			margin-top: 0;
			margin-bottom: 0;
			font-size: 16px;
			width: 730px;
		}
	
		.title-img {
			height: 24px;
			width: 106px;
			display: block;
			margin: auto;
			margin-bottom: 30px;
		}
	
		.password-img {
			width: 106px;
			height: 106px;
			display: block;
			margin: 60px auto;
		}
	
		.border {
			height: 1px;
			width: 100%;
			background-color: #343434;
			transform: scaleY(0.5);
		}
	
		.main-bold {
			font-weight: bold;
			margin-top: 30px;
			font-size: 34px;
		}
	
		.main-normal {
			text-align: center;
			color: #343434;
			margin-top: 30px;
			font-size: 24px;
		}
	
		.code {
			text-align: center;
			font-weight: bold;
			margin: 30px 0;
			font-size: 40px;
		}
	
		.content {
			color: #343434;
			font-size: 22px;
		}
	
		.footer {
			font-size: 15px;
			color: #808080;
			text-align: center;
			margin: 30px 0;
		}
	
		.logo-img {
			height: 60px;
			width: 187px;
		}
	
		@media screen and (max-width: 720px) {
			.main {
				margin: auto;
				margin-top: 0;
				margin-bottom: 0;
				font-size: 12px;
				width: 100%
			}
	
			.title-img {
				height: 12px;
				width: 53px;
				display: block;
				margin: auto;
				margin-bottom: 15px;
			}
	
			.password-img {
				width: 53px;
				height: 53px;
				display: block;
				margin: 30px auto;
			}
	
			.border {
				height: 1px;
				width: 100%;
				background-color: #343434;
				transform: scaleY(0.5);
			}
	
			.main-bold {
				font-weight: bold;
				margin-top: 15px;
				font-size: 18px;
			}
	
			.main-normal {
				text-align: center;
				color: #343434;
				margin-top: 15px;
				font-size: 12px;
			}
	
			.code {
				text-align: center;
				font-weight: bold;
				margin: 15px 0;
				font-size: 24px;
			}
	
			.content {
				color: #343434;
				font-size: 12px;
			}
	
			.footer {
				font-size: 7px;
				color: #808080;
				text-align: center;
				margin: 15px 0;
			}
	
			.logo-img {
				height: 30px;
				width: 93px;
			}
		}
	</style>
	<body>
	<div class="main">
		<img
				class="title-img"
				src="https://static.pivotstudio.cn/husthole/res/husthole.svg"
				alt=""
		/>
		<div class="border"></div>
		<img
				class="password-img"
				src="https://static.pivotstudio.cn/husthole/res/verification.svg"
				alt=""
		/>
		<div class="main-bold">请验证邮箱。</div>
		<div class="main-normal">您的验证码是:</div>
		<div class="code">VerifyCodePlace</div>
		<div class="content">
			<div>如果不是您本人操作, 请忽略此邮件。</div>
			<div style="margin: 12px 0;">祝好！</div>
			<div style="margin-bottom: 20px;">Pivot Studio团队-pivot chat项目组</div>
			<div class="border"></div>
		</div>
		<div class="footer">
			<div class="intro">
			</div>
			<div class="intro">
			</div>
			<div class="intro">联系我们：husthole@pivotstudio.cn</div>
		</div>
	</div>
	</body>
	</html>`
)

const (
	CHAT_CODE_PREFIX = "CHAT_CODE_PREFIX"
)

func init() {
	d = gomail.NewDialer(
		conf.C.EmailServer.Host,
		conf.C.EmailServer.Port,
		conf.C.EmailServer.Email,
		conf.C.EmailServer.Password,
	)
}

// 生成验证码
func CreatCode() (code string) {
	rand.Seed(time.Now().Unix())
	code = fmt.Sprintf("%6v", rand.Intn(600000))
	return
}

// 发送验证码
func SendEmail(ctx context.Context, email string, captcha string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", conf.C.EmailServer.Email)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "邮箱验证")
	content := strings.Replace(emailContent, "VerifyCodePlace", captcha, -1)
	m.SetBody("text/html", content)
	err = d.DialAndSend(m)
	if err != nil {
		logrus.Error("[SendEmail] send to email:%s err:%+v", email, err)
		return err
	}
	return nil
}

// 将验证码存入redis
func CaptchaLogic(ctx *gin.Context, code, email string) {
	codeKey := CHAT_CODE_PREFIX + email
	dao.Cache.Set(ctx, codeKey, code, time.Minute*5) //存入redis 有效5min
}

// 比较验证码
func CaptchaCheck(ctx *gin.Context, input string, email string) bool {
	codeKey := CHAT_CODE_PREFIX + email
	code := dao.Cache.Get(ctx, codeKey).String() //对比验证码是否一致
	return code == input
}
