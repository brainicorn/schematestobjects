package artist

import (
	"github.com/brainicorn/schematestobjects/media"
)

type Artist interface {
	IsBand() bool
}

// @jsonSchema(
//		additionalProperties=true
//		,minProperties=1
//		,maxProperties=20
// )
type Individual struct {
	// @jsonSchema(required=true)
	FirstName   string            `json:"firstName"`
	LastName    string            `json:"lastName,omitempty"`
	Group       bool              `json:"isband,omitempty"`
	SocialMedia media.SocialMedia `json:"social,omitempty"`
}

func (i Individual) IsBand() bool { return false }
