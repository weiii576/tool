package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weiii576/tool/api/user"
)

func SetRouter(server *gin.Engine) {
	user.NewUserRouter(server)
}
