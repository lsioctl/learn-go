package service

import (
	"testing"

	"github.com/lsioctl/rest-layered-poc/fixtures"
	"github.com/lsioctl/rest-layered-poc/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAlbumByID(t *testing.T) {
	testAlbum := fixtures.Album1

	mockDal := mocks.NewDal(t)
	mockDal.EXPECT().GetAlbumByID("1").Return(testAlbum, nil).Once()

	album, _ := mockDal.GetAlbumByID("1")
	assert.Equal(t, testAlbum, album)
}

func TestGetAlbumList(t *testing.T) {
	testAlbumList := fixtures.AlbumList

	mockDal := mocks.NewDal(t)

	mockDal.EXPECT().GetAlbumList().Return(testAlbumList).Once()

	albumList := mockDal.GetAlbumList()
	assert.Equal(t, testAlbumList, albumList)
}
