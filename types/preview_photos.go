package types

type PreviewPhotos struct {
	ID        string `json:"id"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	BlurHash  string `json:"blur_hash"`
	AssetType string `json:"asset_type"`
}
