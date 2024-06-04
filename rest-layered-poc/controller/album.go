package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lsioctl/rest-layered-poc/dto"
	"github.com/lsioctl/rest-layered-poc/service"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbumList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.GetAlbumList())
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, err := service.GetAlbumByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

	c.IndentedJSON(http.StatusOK, album)

}

// postAlbum adds an album from JSON received in the request body.
func PostAlbum(c *gin.Context) {
	var newAlbum dto.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	service.CreateAlbum(newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
