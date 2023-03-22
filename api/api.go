package api

import (
	"github.com/Pivot-Studio/pivot-chat/cors"
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func useRouter(r *gin.Engine) {
	cors_hander := cors.Cors()
	api := r.Group("/api", cors_hander)
	{
		user := api.Group("/user")
		{
			user.POST("/register", Register)
			user.GET("/email", Email)
			user.POST("/chgPwd", ChgPwd)
			user.POST("/login", Login)
			user.GET("/findUserById", FindUserById)
			user.GET("/mygroups", GetMyGroups)
			user.GET("/myjoinedgroups", GetMyJoinedGroups)
		}
		group := api.Group("/group")
		{
			group.GET("/getMembersbyGroupId", GetMembersByGroupId)
			group.GET("/sync", Sync)
			group.POST("/create", CreateGroup)
		}

	}
	r.GET("/ws", wsHandler)
}

func init() {
	Engine = gin.Default()
	useRouter(Engine)
}
