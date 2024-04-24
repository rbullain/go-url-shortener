package api

import (
	"github.com/gin-gonic/gin"
	"go-url-shortener/internal/api/controller"
	"net/http"
)

type Application struct {
	shortenerController controller.ShortenerController
}

func NewShortenerApplication(shortenerController controller.ShortenerController) *Application {
	return &Application{
		shortenerController: shortenerController,
	}
}

func (api *Application) ShortenUrl(ctx *gin.Context) {
	urlDTO, err := api.shortenerController.ShortenUrl(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, urlDTO)
	}
}

func (api *Application) ExpandUrl(ctx *gin.Context) {
	err := api.shortenerController.ExpandUrl(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
