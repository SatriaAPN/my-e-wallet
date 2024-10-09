package handlerrest

import (
	"net/http"

	core "github.com/SatriaAPN/my-e-wallet/backend/pkg/core"
	service "github.com/SatriaAPN/my-e-wallet/backend/pkg/service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
}

type userHandler struct {
	userUsecase service.UserUsecase
}

type UserHandlerConfig struct {
	UserUsecase service.UserUsecase
}

func NewUserHandler(config UserHandlerConfig) UserHandler {
	uh := userHandler{}

	if config.UserUsecase != nil {
		uh.userUsecase = config.UserUsecase
	}

	return &uh
}

func (uh *userHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	u := core.RestCreateUserRequest{}
	var err error

	if err = c.ShouldBindJSON(&u); err != nil {
		c.Error(err)
		return
	}

	u2 := service.CreateUserRequest{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}

	u3, err := uh.userUsecase.CreateUser(ctx, u2)

	if err != nil {
		c.Error(err)
		return
	}

	u4 := core.RestCreateUserResponse{
		Name:  u3.Name,
		Email: u3.Email,
	}

	c.JSON(http.StatusOK, core.Response{Data: u4})
}

func (uh *userHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	u := core.RestLoginUserRequest{}

	if err := c.ShouldBindJSON(&u); err != nil {
		c.Error(err)
		return
	}

	u2 := service.LoginUserRequest{
		Email:    u.Email,
		Password: u.Password,
	}

	u3, err := uh.userUsecase.LoginUser(ctx, u2)

	if err != nil {
		c.Error(err)
		return
	}

	u4 := core.RestLoginUserResponse{
		Token: u3.Token,
	}

	c.JSON(http.StatusOK, core.Response{Data: u4})
}
