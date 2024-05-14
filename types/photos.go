package types

type Photos struct {
	ID             string `json:"id"`
	Slug           string `json:"slug"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	BlurHash       string `json:"blur_hash"`
	AssetType      string `json:"asset_type"`
	PromotedAt     string `json:"promoted_at"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Color          string `json:"color"`
	Description    string `json:"description"`
	AltDescription string `json:"alt_description"`
	Likes          int    `json:"likes"`
	LikedByUser    bool   `json:"liked_by_user"`
	Sponsorship    string `json:"sponsorship"`
}
