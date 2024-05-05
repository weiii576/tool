package user

import (
	"github.com/gin-gonic/gin"
	"github.com/weiii576/tool/storage"
)

type UserController struct {
	store *storage.PostgresStore
}

func NewUserController(store *storage.PostgresStore) *UserController {
	return &UserController{
		store: store,
	}
}

func (uc *UserController) handleCreateUser(c *gin.Context) {
}
