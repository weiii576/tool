package user

import "github.com/gin-gonic/gin"

func NewUserRouter(server *gin.Engine) {
	uc := NewUserController()

	group := server.Group("/user")

	group.POST("", uc.handleCreateUser)
}
