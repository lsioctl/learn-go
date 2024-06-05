package dal

import (
	"errors"
	"fmt"

	"github.com/lsioctl/rest-layered-poc/dto"
)

// albums slice to seed record album data.
var albums = []dto.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func (dal *dal) GetAlbumList() []dto.Album {
	fmt.Println(dal.db)
	return albums
}

func (*dal) GetAlbumByID(id string) (dto.Album, error) {
	if id == "" {
		return dto.Album{}, errors.New("empty id")
	}

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			return a, nil
		}
	}

	return dto.Album{}, errors.New("not found")
}

// TODO: do not like the prefix
func (*dal) CreateAlbum(album dto.Album) {
	// Add the new album to the slice.
	albums = append(albums, album)
}
