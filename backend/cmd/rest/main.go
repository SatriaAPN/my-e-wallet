package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SatriaAPN/my-e-wallet/backend/pkg/database"

	"github.com/SatriaAPN/my-e-wallet/backend/pkg/core"
	"github.com/SatriaAPN/my-e-wallet/backend/pkg/middleware"

	handlerrest "github.com/SatriaAPN/my-e-wallet/backend/pkg/api/rest"

	"github.com/SatriaAPN/my-e-wallet/backend/pkg/config"
	"github.com/SatriaAPN/my-e-wallet/backend/pkg/repository"
	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/SatriaAPN/my-e-wallet/backend/pkg/service"

	"gorm.io/gorm"
)

func main() {
	InitServer()
}

type Router interface {
	GetRouter() *gin.Engine
}

type router struct {
	router      *gin.Engine
	userHandler handlerrest.UserHandler
}

type RouterConfig struct {
	UserHandler handlerrest.UserHandler
}

func NewRouter(config RouterConfig) Router {
	return &router{
		router:      gin.New(),
		userHandler: config.UserHandler,
	}
}

func (rb *router) GetRouter() *gin.Engine {
	rb.buildMiddleware()
	rb.buildEndpointHandler()

	return rb.router
}

func (rb *router) buildEndpointHandler() {
	rb.router.POST("/users", rb.userHandler.CreateUser)
	rb.router.POST("/login", rb.userHandler.Login)
}

func (rb *router) buildMiddleware() {
	rb.router.Use(middleware.SetRequestId())

	rb.router.Use(middleware.Logger())

	rb.router.Use(middleware.GlobalErrorHandler())

	rb.router.Use(middleware.HttpRequestTimeout())
}

func NewServer() *http.Server {

	r := NewRouter(initHandler()).GetRouter()

	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return &s
}

func InitServer() {
	config.InitEnvReader()

	srv := NewServer()

	GracefulShutdown(srv)
}

func initHandler() RouterConfig {
	db := database.GetInstance()

	nrc := RouterConfig{
		UserHandler: initUserHandler(db),
	}

	return nrc
}

func initUserHandler(db *gorm.DB) handlerrest.UserHandler {
	urc := repository.UserRepositoryConfig{
		Db: db,
	}
	ur := repository.NewUserRepository(urc)

	uuc := service.UserUsecaseConfig{
		UserRepository:       ur,
		PasswordHasher:       core.GetPasswordHasher(),
		AuthTokenGenerator:   core.GetAuthTokenGenerator(),
		RandomTokenGenerator: core.GetRandomTokenGenerator(),
	}
	uu := service.NewUserUsecase(uuc)

	uhc := handlerrest.UserHandlerConfig{
		UserUsecase: uu,
	}
	uh := handlerrest.NewUserHandler(uhc)

	return uh
}

func GracefulShutdown(srv *http.Server) {
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
