package main

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/lsioctl/rest-layered-poc/controller"
	"github.com/lsioctl/rest-layered-poc/dal"
	"github.com/lsioctl/rest-layered-poc/service"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Printf("num CPUs: %d\n", numCPU)
	store := dal.New()
	service := service.New(store)
	controller := controller.New(service)

	router := gin.Default()

	router.GET("/albums", controller.GetAlbumList)
	router.GET("/albums/:id", controller.GetAlbumByID)
	router.POST("/albums", controller.PostAlbum)

	router.Run("localhost:8083")
}
