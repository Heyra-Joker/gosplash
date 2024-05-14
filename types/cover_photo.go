package types

type CoverPhoto struct {
	ID             string `json:"id"`
	Slug           string `json:"slug"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	PromotedAt     string `json:"promoted_at"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Color          string `json:"color"`
	BlurHash       string `json:"blur_hash"`
	Description    string `json:"description"`
	AltDescription string `json:"alt_description"`
	Likes          int    `json:"likes"`
	LikedByUser    bool   `json:"liked_by_user"`
	Sponsorship    string `json:"sponsorship"`
	AssetType      string `json:"asset_type"`
	Premium        bool   `json:"premium"`
	Plus           bool   `json:"plus"`
}
