package user

import (
	"github.com/gin-gonic/gin"
	"github.com/weiii576/tool/configs"
	"github.com/weiii576/tool/storage"
)

func NewUserRouter(server *gin.Engine, store *storage.PostgresStore, env *configs.Env) {
	uc := NewUserController(store, env)

	group := server.Group("/user")

	group.POST("", uc.handleCreateUser)
}
