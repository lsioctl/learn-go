package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lsioctl/rest-layered-poc/service"
)

type controller struct {
	service service.Service
}

func New(service service.Service) *controller {
	return &controller{
		service: service,
	}
}

type Controller interface {
	GetAlbumList(c *gin.Context)
	GetAlbumByID(c *gin.Context)
	PostAlbum(c *gin.Context)
}
