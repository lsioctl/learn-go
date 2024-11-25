package service

import (
	"fmt"
	"runtime"
	"time"

	"github.com/lsioctl/rest-layered-poc/dto"
)

// TODO: There is no value to the layer here, try a bit more realistic
// example
func (service *service) GetAlbumList() []dto.Album {
	numGo := runtime.NumGoroutine()
	fmt.Printf("num goroutines: %d\n", numGo)
	fmt.Println("starting to sleep")
	time.Sleep(8 * time.Second)
	fmt.Println("ending sleep")
	return service.store.GetAlbumList()
}

func (service *service) GetAlbumByID(id string) (dto.Album, error) {
	album, err := service.store.GetAlbumByID(id)

	return album, err
}

func (service *service) CreateAlbum(album dto.Album) {
	service.store.CreateAlbum(album)
}
