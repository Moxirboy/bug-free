package http

import (
	config2 "bug-free/internal/config"
	"bug-free/internal/http/handler"
	"bug-free/internal/service/repo"
	usecase "bug-free/internal/service/usecase"
	postgres "bug-free/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	config := config2.Configuration()
	db, err := postgres.NewPostgresConfig(config)
	if err != nil {
		panic("no database connection")
	}
	newPostgres := repo.NewPostgres(db)
	service := usecase.NewUser(newPostgres)
	handle := handler.NewHandler(service)
	return handle.InitRouter()
}
