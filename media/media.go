package media

// MediaURL is a url to an image or something.
// These are useful for linking to images or to social media sites.
//
// @jsonSchema(format="ipv4")
type MediaURL string

type SocialMedia struct {
	Twitter MediaURL `json:"twitter,omitempty"`

	// Facebook is... well.. Facebook.
	Facebook MediaURL `json:"facebook,omitempty"`

	// Instagram picture this.
	Instagram MediaURL `json:"instagram,omitempty"`

	// @jsonSchema(format="email")
	Email string `json:"email,omitempty"`

	// @jsonSchema(minLength=2, maxLength=200, pattern="[A-Za-z0-9]")
	ChatMention string `json:"mention,omitempty"`
}
