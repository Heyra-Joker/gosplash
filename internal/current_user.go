package internal

import (
	"fmt"
	"time"
)

// CurrentUser Group
type CurrentUser struct {
	Client *Client
}

// MeRequest is empty
type MeRequest struct {
}

// MeReply get-the-users-profile response
type MeReply struct {
	ID                string `json:"id"`
	UpdatedAt         string `json:"updated_at"`
	Username          string `json:"username"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	TwitterUsername   string `json:"twitter_username"`
	PortfolioURL      string `json:"portfolio_url"`
	Bio               string `json:"bio"`
	Location          string `json:"location"`
	TotalLikes        int    `json:"total_likes"`
	TotalPhotos       int    `json:"total_photos"`
	TotalCollections  int    `json:"total_collections"`
	FollowedByUser    bool   `json:"followed_by_user"`
	Downloads         int    `json:"downloads"`
	UploadsRemaining  int    `json:"uploads_remaining"`
	InstagramUsername string `json:"instagram_username"`
	Email             string `json:"email"`
	Links             struct {
		Self      string `json:"self"`
		HTML      string `json:"html"`
		Photos    string `json:"photos"`
		Likes     string `json:"likes"`
		Portfolio string `json:"portfolio"`
	} `json:"links"`
}

// API get-the-users-profile URL
func (c MeRequest) API() string {
	return "https://api.unsplash.com/me"
}

// Params get-the-users-profile params is empty
func (c MeRequest) Params() map[string]string {
	return map[string]string{}
}

// Me Get the user’s profile
// document: https://unsplash.com/documentation#get-the-users-profile
func (c *CurrentUser) Me(token string) (reply *MeReply, headers *ResponseHeaders, err error) {
	reqHeaders := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	headers, err = c.Client.request("GET", MeRequest{}, &reply, reqHeaders)
	return
}

// UpdateMeRequest update me request params
type UpdateMeRequest struct {
	UserName          string `json:"user_name"`
	FirstName         string `json:"first_name"` // Must Param
	LastName          string `json:"last_name"`
	Email             string `json:"email"` // Must Param
	URL               string `json:"url"`
	Location          string `json:"location"`
	BIO               string `json:"bio"`
	InstagramUsername string `json:"instagram_username"`
}

// UpdateMeReply update me api response
type UpdateMeReply struct {
	ID              string    `json:"id"`
	UpdatedAt       time.Time `json:"updated_at"`
	Username        string    `json:"username"`
	Name            string    `json:"name"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	TwitterUsername string    `json:"twitter_username"`
	PortfolioURL    string    `json:"portfolio_url"`
	Bio             string    `json:"bio"`
	Location        string    `json:"location"`
	Links           struct {
		Self      string `json:"self"`
		HTML      string `json:"html"`
		Photos    string `json:"photos"`
		Likes     string `json:"likes"`
		Portfolio string `json:"portfolio"`
		Following string `json:"following"`
		Followers string `json:"followers"`
	} `json:"links"`
	ProfileImage struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"profile_image"`
	InstagramUsername          string `json:"instagram_username"`
	TotalCollections           int    `json:"total_collections"`
	TotalLikes                 int    `json:"total_likes"`
	TotalPhotos                int    `json:"total_photos"`
	TotalPromotedPhotos        int    `json:"total_promoted_photos"`
	TotalIllustrations         int    `json:"total_illustrations"`
	TotalPromotedIllustrations int    `json:"total_promoted_illustrations"`
	AcceptedTos                bool   `json:"accepted_tos"`
	ForHire                    bool   `json:"for_hire"`
	Social                     struct {
		InstagramUsername string `json:"instagram_username"`
		PortfolioURL      string `json:"portfolio_url"`
		TwitterUsername   string `json:"twitter_username"`
		PaypalEmail       string `json:"paypal_email"`
	} `json:"social"`
	FollowedByUser bool     `json:"followed_by_user"`
	Photos         []string `json:"photos"`
	Badge          string   `json:"badge"`
	Tags           struct {
		Custom     []string `json:"custom"`
		Aggregated []string `json:"aggregated"`
	} `json:"tags"`
	FollowersCount int  `json:"followers_count"`
	FollowingCount int  `json:"following_count"`
	AllowMessages  bool `json:"allow_messages"`
	NumericID      int  `json:"numeric_id"`
	Downloads      int  `json:"downloads"`
	Meta           struct {
		Index bool `json:"index"`
	} `json:"meta"`
	UID                          string `json:"uid"`
	Confirmed                    bool   `json:"confirmed"`
	UploadsRemaining             int    `json:"uploads_remaining"`
	UnlimitedUploads             bool   `json:"unlimited_uploads"`
	Email                        string `json:"email"`
	DmcaVerification             string `json:"dmca_verification"`
	UnreadInAppNotifications     bool   `json:"unread_in_app_notifications"`
	UnreadHighlightNotifications bool   `json:"unread_highlight_notifications"`
}

// API Update the current user’s profile url
func (r UpdateMeRequest) API() string {
	return "https://api.unsplash.com/me"
}

// Params Update the current user’s profile params is empty
func (r UpdateMeRequest) Params() map[string]string {
	return map[string]string{}
}

// UpdateMe Update the current user’s profile
// Note: This action requires the write_user scope. Without it, it will return a 403 Forbidden response.
// document: https://unsplash.com/documentation#update-the-current-users-profile
func (c *CurrentUser) UpdateMe(token string, req Req) (reply *UpdateMeReply, headers *ResponseHeaders, err error) {
	reqHeaders := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	headers, err = c.Client.request("PUT", req, &reply, reqHeaders)
	return
}
