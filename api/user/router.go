package user

import (
	"github.com/gin-gonic/gin"
	"github.com/weiii576/tool/storage"
)

func NewUserRouter(server *gin.Engine, store *storage.PostgresStore) {
	uc := NewUserController(store)

	group := server.Group("/user")

	group.POST("", uc.handleCreateUser)
}
