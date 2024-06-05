package service

import "github.com/lsioctl/rest-layered-poc/dto"

// TODO: There is no value to the layer here, try a bit more realistic
// example
func (service *service) GetAlbumList() []dto.Album {
	return service.store.GetAlbumList()
}

func (service *service) GetAlbumByID(id string) (dto.Album, error) {
	album, err := service.store.GetAlbumByID(id)

	return album, err
}

func (service *service) CreateAlbum(album dto.Album) {
	service.store.CreateAlbum(album)
}
