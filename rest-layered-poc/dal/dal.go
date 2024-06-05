package dal

import (
	"fmt"

	"github.com/lsioctl/rest-layered-poc/dto"
)

// For now simulate ownership of a DB pool
type dal struct {
	db int
}

// now that we own a db, we have to instanciate
func New() *dal {
	fmt.Println("new called")
	return &dal{
		db: 3,
	}
}

// and make methods of the previously free function
// so they can access the DB
type Dal interface {
	GetAlbumList() []dto.Album
	GetAlbumByID(id string) (dto.Album, error)
	CreateAlbum(album dto.Album)
}
