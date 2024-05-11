package user

import (
	"github.com/gin-gonic/gin"
	"github.com/weiii576/tool/configs"
	"gorm.io/gorm"
)

func NewUserRouter(server *gin.Engine, store *gorm.DB, env *configs.Env) {
	us := NewUserStorage(store)
	uc := NewUserController(us, env)

	group := server.Group("/user")

	group.POST("", uc.handleCreateUser)
	group.GET("", uc.handleGetUsers)
}
