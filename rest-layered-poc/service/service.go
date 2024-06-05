package service

import (
	"github.com/lsioctl/rest-layered-poc/dal"
	"github.com/lsioctl/rest-layered-poc/dto"
)

type service struct {
	store dal.Dal
}

func New(store dal.Dal) *service {
	return &service{
		store: store,
	}
}

type Service interface {
	GetAlbumList() []dto.Album
	GetAlbumByID(id string) (dto.Album, error)
	CreateAlbum(album dto.Album)
}
