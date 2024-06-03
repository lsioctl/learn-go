package service

import (
	"github.com/lsioctl/rest-layered-poc/dal"
)

// TODO: There is no value to the layer here, try a bit more realistic
// example
func GetAlbumList() []dal.Album {
	return dal.GetAlbumList()
}

func GetAlbumByID(id string) (dal.Album, error) {
	album, err := dal.GetAlbumByID(id)

	return album, err
}
