package api

import "github.com/gin-gonic/gin"

var Engine *gin.Engine

func useRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/register", Register)
			user.GET("/email", Email)
			user.POST("/chgPwd", ChgPwd)
			user.POST("/login", Login)
			user.GET("/getMembersbyGroupId", GetMembersByGroupId)
			user.GET("/findUserById", FindUserById)
		}
	}
	r.GET("/ws", wsHandler)
}

func init() {
	Engine = gin.Default()
	useRouter(Engine)
}
