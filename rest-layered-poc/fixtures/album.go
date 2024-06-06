package fixtures

import "github.com/lsioctl/rest-layered-poc/dto"

var Album1 = dto.Album{
	ID:     "1",
	Artist: "J",
	Title:  "T",
	Price:  900,
}

var Album2 = dto.Album{
	ID:     "2",
	Artist: "J",
	Title:  "T",
	Price:  900,
}

// Funny (or not) in Go [2]dto.Album is a different type from
// []dto.Album (slice) or [x]dto.Album (array of different size)
// Even if I don't care copying the array, slice are needed
// if I want to pass them as function arguments
// To rember, rhs belows:
// * create the array with two albums
// * then a slice that reference it
var AlbumList = []dto.Album{Album1, Album2}
