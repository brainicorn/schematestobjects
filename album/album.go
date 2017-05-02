package album

import (
	"time"

	"github.com/brainicorn/schematestobjects/artist"
	"github.com/brainicorn/schematestobjects/media"
	"github.com/brainicorn/to"
)

type Song struct {
	// @jsonSchema(required=true)
	SongTitle string `json:"title"`
}

// @jsonSchema(definition="merchandise")
type ForSale struct {
	// @jsonSchema(multipleOf=5, required=true)
	Price float64

	// @jsonSchema(
	//		minimum=0
	//		,maximum=5
	//		,exclusiveMinimum=true
	//		,exclusiveMaximum=false
	// )
	Rating int `json:"rating,omitempty"`
}

// Album is a collection of recorded songs.
// People listen to albums, although nowadays they really just buy individual songs.
//
// @jsonSchema(
//		id="http://github.com/schemas/album"
//		,title="An Album."
//		,description="A thing we used to play with a player and now we get from the air"
// )
type Album struct {
	ForSale

	//Publisher artist.Individual

	// @jsonSchema(required=true)
	AlbumTitle to.Title

	// @jsonSchema(default="mp3")
	Format string `json:"format,omitempty"`

	// Artist is a person or band.
	// The term artist is used loosely here as any  ol' shmoe can record what they call songs.
	//
	// @jsonSchema(required=true,
	// 	oneOf=["github.com/brainicorn/schematestobjects/artist/Band"
	//	,"github.com/brainicorn/schematestobjects/artist/Individual"])
	Artist artist.Artist `json:"artist"`

	// @jsonSchema(
	//		required=true, maxItems=50, minItems=1, additionalItems=false, uniqueItems=true
	// )
	Songs []*Song `json:"tracks"`

	// Cover is the image that would have been on the cover if it was a tangible good.
	// everyone loves covers.
	//
	// @jsonSchema(required=true)
	Cover media.MediaURL

	RecordedAt string `json:"-"`

	// @jsonSchema(required=true)
	ReleaseDate time.Time
}
