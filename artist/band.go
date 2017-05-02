package artist

import (
	"github.com/brainicorn/schematestobjects/media"
)

type TourRider map[string]string

// @jsonSchema(minimum=1)
type TVAppearances int

type Band struct {
	// @jsonSchema(required=true)
	Name        string       `json:"name"`
	Members     []Individual `json:"members,omitempty"`
	Group       bool         `json:"isband,omitempty"`
	SocialMedia media.SocialMedia

	// @jsonSchema(title="Rider", description="A list of stuff a band requires to play somewhere")
	Rider TourRider

	TV TVAppearances `json:"tv,omitempty"`
}

func (b Band) IsBand() bool { return true }
