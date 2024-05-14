package types

type User struct {
	ID                           string `json:"id"`
	UpdatedAt                    string `json:"updated_at"`
	Username                     string `json:"username"`
	Name                         string `json:"name"`
	FirstName                    string `json:"first_name"`
	LastName                     string `json:"last_name"`
	TwitterUsername              string `json:"twitter_username"`
	PortfolioURL                 string `json:"portfolio_url"`
	Bio                          string `json:"bio"`
	Location                     string `json:"location"`
	InstagramUsername            string `json:"instagram_username"`
	TotalCollections             int    `json:"total_collections"`
	TotalLikes                   int    `json:"total_likes"`
	TotalPhotos                  int    `json:"total_photos"`
	TotalPromotedPhotos          int    `json:"total_promoted_photos"`
	TotalIllustrations           int    `json:"total_illustrations"`
	TotalPromotedIllustrations   int    `json:"total_promoted_illustrations"`
	AcceptedTos                  bool   `json:"accepted_tos"`
	ForHire                      bool   `json:"for_hire"`
	FollowedByUser               bool   `json:"followed_by_user"`
	Badge                        string `json:"badge"`
	FollowersCount               int    `json:"followers_count"`
	FollowingCount               int    `json:"following_count"`
	AllowMessages                bool   `json:"allow_messages"`
	NumericID                    int    `json:"numeric_id"`
	Downloads                    int    `json:"downloads"`
	UID                          string `json:"uid"`
	Confirmed                    bool   `json:"confirmed"`
	UploadsRemaining             int    `json:"uploads_remaining"`
	UnlimitedUploads             bool   `json:"unlimited_uploads"`
	Email                        string `json:"email"`
	DmcaVerification             string `json:"dmca_verification"`
	UnreadInAppNotifications     bool   `json:"unread_in_app_notifications"`
	UnreadHighlightNotifications bool   `json:"unread_highlight_notifications"`
}
