package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weiii576/tool/api/user"
	"github.com/weiii576/tool/configs"
	"github.com/weiii576/tool/storage"
)

func SetRouter(server *gin.Engine, store *storage.PostgresStore, env *configs.Env) {
	user.NewUserRouter(server, store, env)
}
