package rest

import (
	"go-template/database"
	handlerrest "go-template/handler/rest"
	"go-template/repository"
	routerrest "go-template/server/rest/router"
	"go-template/share/general/config"
	"go-template/share/general/util"
	utilhttp "go-template/share/http/util"
	"go-template/usecase"
	"net/http"

	"gorm.io/gorm"
)

func InitServer() {
	config.InitEnvReader()

	srv := NewServer()

	utilhttp.GracefulShutdown(srv)
}

func NewServer() *http.Server {

	r := routerrest.NewRouter(initHandler()).GetRouter()

	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return &s
}

func initHandler() routerrest.RouterConfig {
	db := database.GetInstance()

	nrc := routerrest.RouterConfig{
		UserHandler: initUserHandler(db),
	}

	return nrc
}

func initUserHandler(db *gorm.DB) handlerrest.UserHandler {
	urc := repository.UserRepositoryConfig{
		Db: db,
	}
	ur := repository.NewUserRepository(urc)

	uuc := usecase.UserUsecaseConfig{
		UserRepository:       ur,
		PasswordHasher:       util.GetPasswordHasher(),
		AuthTokenGenerator:   util.GetAuthTokenGenerator(),
		RandomTokenGenerator: util.GetRandomTokenGenerator(),
	}
	uu := usecase.NewUserUsecase(uuc)

	uhc := handlerrest.UserHandlerConfig{
		UserUsecase: uu,
	}
	uh := handlerrest.NewUserHandler(uhc)

	return uh
}
