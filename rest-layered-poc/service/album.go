package service

import (
	"github.com/lsioctl/rest-layered-poc/dal"
	"github.com/lsioctl/rest-layered-poc/dto"
)

// TODO: There is no value to the layer here, try a bit more realistic
// example
func GetAlbumList() []dto.Album {
	return dal.GetAlbumList()
}

func GetAlbumByID(id string) (dto.Album, error) {
	album, err := dal.GetAlbumByID(id)

	return album, err
}

func CreateAlbum(album dto.Album) {
	dal.CreateAlbum(album)
}
