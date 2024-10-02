package handlerrest

import (
	dtousecase "go-template/dto/general/usecase"
	dtohttp "go-template/dto/http"
	"go-template/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
	ForgetPassword(c *gin.Context)
	ResetPassword(c *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

type UserHandlerConfig struct {
	UserUsecase usecase.UserUsecase
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
	u := dtohttp.CreateUserRequest{}
	var err error

	if err = c.ShouldBindJSON(&u); err != nil {
		c.Error(err)
		return
	}

	u2 := dtousecase.CreateUserRequest{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}

	u3, err := uh.userUsecase.CreateUser(ctx, u2)

	if err != nil {
		c.Error(err)
		return
	}

	u4 := dtohttp.CreateUserResponse{
		Name:  u3.Name,
		Email: u3.Email,
	}

	c.JSON(http.StatusOK, dtohttp.Response{Data: u4})
}

func (uh *userHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	u := dtohttp.LoginUserRequest{}

	if err := c.ShouldBindJSON(&u); err != nil {
		c.Error(err)
		return
	}

	u2 := dtousecase.LoginUserRequest{
		Email:    u.Email,
		Password: u.Password,
	}

	u3, err := uh.userUsecase.LoginUser(ctx, u2)

	if err != nil {
		c.Error(err)
		return
	}

	u4 := dtohttp.LoginUserResponse{
		Token: u3.Token,
	}

	c.JSON(http.StatusOK, dtohttp.Response{Data: u4})
}

func (uh *userHandler) ForgetPassword(c *gin.Context) {
	ctx := c.Request.Context()
	res := dtohttp.ForgetPasswordResponse{}
	body := dtohttp.ForgetPasswordRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		return
	}

	ub := dtousecase.ForgetPasswordRequest{
		Email: body.Email,
	}

	data, err := uh.userUsecase.ForgetPassword(ctx, ub)

	if err != nil {
		c.Error(err)
		return
	}

	res.Token = data.Token
	res.ExpiredAt = data.ExpiredAt

	c.JSON(http.StatusOK, dtohttp.Response{Data: res})
}

func (uh *userHandler) ResetPassword(c *gin.Context) {
	ctx := c.Request.Context()

	body := dtohttp.ResetPasswordRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		return
	}

	body2 := dtousecase.ResetPasswordRequest{
		Email:       body.Email,
		NewPassword: body.NewPassword,
		Token:       body.Token,
	}

	err := uh.userUsecase.ResetPassword(ctx, body2)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtohttp.Response{Message: "success change password"})
}
