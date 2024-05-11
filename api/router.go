package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weiii576/tool/api/user"
	"github.com/weiii576/tool/configs"
	"gorm.io/gorm"
)

func SetRouter(server *gin.Engine, store *gorm.DB, env *configs.Env) {
	user.NewUserRouter(server, store, env)
}
