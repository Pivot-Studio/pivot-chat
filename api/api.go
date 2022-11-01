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
			user.GET("/findUserById", FindUserById)
			user.GET("/mygroups", GetMyGroups)
		}
		group := api.Group("/group")
		{
			group.GET("/getMembersbyGroupId", GetMembersByGroupId)
			group.GET("/sync", Sync)

		}

	}
	r.GET("/ws", wsHandler)
}

func init() {
	Engine = gin.Default()
	useRouter(Engine)
}
