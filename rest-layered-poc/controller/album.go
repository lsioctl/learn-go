package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lsioctl/rest-layered-poc/dto"
)

// getAlbums responds with the list of all albums as JSON.
func (controller *controller) GetAlbumList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, controller.service.GetAlbumList())
}

func printContext(ctx context.Context) {
	fmt.Println(ctx)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (controller *controller) GetAlbumByID(c *gin.Context) {
	printContext(c)
	id := c.Param("id")

	album, err := controller.service.GetAlbumByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

	c.IndentedJSON(http.StatusOK, album)

}

// postAlbum adds an album from JSON received in the request body.
func (controller *controller) PostAlbum(c *gin.Context) {
	var newAlbum dto.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	controller.service.CreateAlbum(newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
