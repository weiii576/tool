package user

import (
	"github.com/gin-gonic/gin"
	"github.com/weiii576/tool/configs"
	"github.com/weiii576/tool/storage"
)

func NewUserRouter(server *gin.Engine, store *storage.PostgresStore, env *configs.Env) {
	us := NewUserStorage(store)
	uc := NewUserController(us, env)

	group := server.Group("/user")

	group.POST("", uc.handleCreateUser)
}
