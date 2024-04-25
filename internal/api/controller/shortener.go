package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-url-shortener/internal/api/dto"
	"go-url-shortener/internal/domain/shortener/service"
	"net/http"
)

type ShortenerController interface {
	ShortenUrl(ctx *gin.Context) (*dto.ShortenUrlResponseDTO, error)
	ExpandUrl(ctx *gin.Context) error
}

type shortenerController struct {
	service service.ShortenerService
}

func NewShortenerController(service service.ShortenerService) ShortenerController {
	return &shortenerController{
		service: service,
	}
}

func (controller *shortenerController) ShortenUrl(ctx *gin.Context) (*dto.ShortenUrlResponseDTO, error) {
	var urlDTO dto.ShortenUrlRequestDTO

	err := ctx.ShouldBindJSON(&urlDTO)
	if err != nil {
		return nil, err
	}

	shortenUrl, err := controller.service.ShortenUrl(urlDTO.Url)
	if err != nil {
		return nil, err
	}

	response := &dto.ShortenUrlResponseDTO{
		Url:            shortenUrl.OriginalUrl,
		ShortUrl:       shortenUrl.Hash,
		ExpirationDate: shortenUrl.ExpirationDate,
	}
	return response, nil
}

func (controller *shortenerController) ExpandUrl(ctx *gin.Context) error {
	url := ctx.Param("url")
	if url == "" {
		return errors.New("url parameter is required")
	}

	expandedUrl, err := controller.service.ExpandUrl(url)
	if err != nil {
		return err
	}

	if expandedUrl == nil {
		return errors.New("url does not exist")
	}

	ctx.Redirect(http.StatusFound, expandedUrl.OriginalUrl)
	return nil
}
