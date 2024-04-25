package main

import (
	"github.com/gin-gonic/gin"
	"go-url-shortener/internal/api"
	"go-url-shortener/internal/api/controller"
	"go-url-shortener/internal/application"
	"go-url-shortener/internal/infra/repository"
)

var (
	shortenerRepository  = repository.NewDatabaseShortenerRepository("admin", "admin", "localhost", "3306", "golang_learning")
	shortenerService     = application.NewSum256HashShortener(shortenerRepository)
	shortenerController  = controller.NewShortenerController(shortenerService)
	shortenerApplication = api.NewShortenerApplication(shortenerController)
)

func main() {
	router := gin.Default()

	router.POST("/", shortenerApplication.ShortenUrl)
	router.GET("/:url", shortenerApplication.ExpandUrl)

	_ = router.Run("localhost:8080")
}
