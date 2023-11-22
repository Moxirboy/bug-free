package http

import (
	"github.com/gin-gonic/gin"
	config2 "golang-project-template/internal/config"
	"golang-project-template/internal/http/handler"
	"golang-project-template/internal/service/repo"
	usecase "golang-project-template/internal/service/usecase"
	postgres "golang-project-template/pkg/postgres"
)

func Router() *gin.Engine {
	config := config2.Configuration()
	db, err := postgres.NewPostgresConfig(config)
	if err != nil {
		panic("no database connection")
	}
	newPostgres := repo.NewPostgres(db)
	service := usecase.NewCrud(newPostgres)
	handle := handler.NewHandler(service)
	return handle.InitRouter()
}
