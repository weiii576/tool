package user

import (
	"github.com/gin-gonic/gin"
	"github.com/weiii576/tool/configs"
	"github.com/weiii576/tool/storage"
)

type UserController struct {
	store *storage.PostgresStore
	env   *configs.Env
}

func NewUserController(store *storage.PostgresStore, env *configs.Env) *UserController {
	return &UserController{
		store: store,
		env:   env,
	}
}

func (uc *UserController) handleCreateUser(c *gin.Context) {
}
