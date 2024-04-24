package main

import (
	"github.com/gin-gonic/gin"
	"go-url-shortener/internal/api"
	"go-url-shortener/internal/api/controller"
	"go-url-shortener/internal/repository"
	"go-url-shortener/internal/service"
)

var (
	shortenerRepository  = repository.NewDatabaseShortenerRepository("admin", "admin", "localhost", "3306", "golang_learning")
	shortenerService     = service.NewShortenerService(shortenerRepository)
	shortenerController  = controller.NewShortenerController(shortenerService)
	shortenerApplication = api.NewShortenerApplication(shortenerController)
)

func main() {
	router := gin.Default()

	router.POST("/", shortenerApplication.ShortenUrl)
	router.GET("/:url", shortenerApplication.ExpandUrl)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
