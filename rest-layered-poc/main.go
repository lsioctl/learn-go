package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lsioctl/rest-layered-poc/controller"
)

func main() {
	router := gin.Default()
	router.GET("/albums", controller.GetAlbumList)
	router.GET("/albums/:id", controller.GetAlbumByID)
	router.POST("/albums", controller.PostAlbum)

	router.Run("localhost:8083")
}
