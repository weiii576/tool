package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/weiii576/tool/configs"
	"github.com/weiii576/tool/internal"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserController interface {
		handleCreateUser(*gin.Context)
		handleGetUsers(*gin.Context)
	}

	userController struct {
		UserStorage UserStorage
		Env         *configs.Env
	}
)

func NewUserController(userStorage UserStorage, env *configs.Env) UserController {
	return &userController{
		UserStorage: userStorage,
		Env:         env,
	}
}

func (uc *userController) handleCreateUser(ctx *gin.Context) {
	var request CreateUserRequest

	if err := ctx.ShouldBind(&request); err != nil {
		res := internal.BuildResponseFailed(MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	_, err := uc.UserStorage.GetUserByEmail(request.Email)
	if err == nil {
		res := internal.BuildResponseFailed(MESSAGE_EMAIL_ALREADY_USED, "", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, internal.BuildInternalServerError())
		return
	}

	request.Password = string(encryptedPassword)

	user := &User{
		ID:       uuid.New(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = uc.UserStorage.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, internal.BuildResponseFailed(MESSAGE_FAILED_CREATE_USER, err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, internal.BuildResponseSuccess(MESSAGE_SUCCESS_CREATE_USER, nil))
}

func (uc *userController) handleGetUsers(ctx *gin.Context) {
	var request internal.PaginationRequest

	if err := ctx.ShouldBind(&request); err != nil {
		res := internal.BuildResponseFailed(MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	users, pagination, err := uc.UserStorage.GetUsers(&request)
	if err != nil {
		res := internal.BuildResponseFailed(MESSAGE_FAILED_GET_USERS, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	paginationResponse := GetUsersResponse{
		Users:      *users,
		Pagination: *pagination,
	}

	ctx.JSON(http.StatusOK, internal.BuildResponseSuccess(MESSAGE_SUCCESS_GET_USERS, paginationResponse))
}
