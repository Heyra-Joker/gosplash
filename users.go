package main

import (
	"github.com/Heyra-Joker/gosplash/types"
)

// Users Groups
type Users struct {
	Client *Client
}

// UserGroupReq .
type UserGroupReq struct {
}

func (p UserGroupReq) API() string {
	return "https://api.unsplash.com/users/"
}

func (p UserGroupReq) Params() map[string]string {
	return nil
}

// UserPublicProfileReply .
type UserPublicProfileReply struct {
	types.User
	Links        types.Links        `json:"links"`
	ProfileImage types.ProfileImage `json:"profile_image"`
	Social       types.Social       `json:"social"`
	Photos       []struct {
		types.Photos
		URLs types.URLs `json:"urls"`
	} `json:"photos"`
	Tags struct {
		Custom []struct {
			types.Custom
			Source struct {
				types.Source
				Ancestry struct {
					Type        types.Type        `json:"type"`
					Category    types.Category    `json:"category"`
					Subcategory types.Subcategory `json:"subcategory"`
				} `json:"ancestry"`
				CoverPhoto struct {
					types.CoverPhoto
					AlternativeSlugs       types.AlternativeSlugs         `json:"alternative_slugs"`
					Breadcrumbs            []types.Breadcrumbs            `json:"breadcrumbs"`
					URLs                   types.URLs                     `json:"urls"`
					Links                  types.Links                    `json:"links"`
					CurrentUserCollections []types.CurrentUserCollections `json:"current_user_collections"`
					TopicSubmissions       types.TopicSubmissions         `json:"topic_submissions"`
					User                   struct {
						types.User
						Links        types.Links        `json:"links"`
						ProfileImage types.ProfileImage `json:"profile_image"`
						Social       types.Social       `json:"social"`
					} `json:"user"`
				} `json:"cover_photo"`
			} `json:"source"`
		} `json:"custom"`
		Aggregated []struct {
			types.Aggregated
			Source struct {
				types.Source
				Ancestry struct {
					Type        types.Type        `json:"type"`
					Category    types.Category    `json:"category"`
					Subcategory types.Subcategory `json:"subcategory"`
				} `json:"ancestry"`
				CoverPhoto struct {
					types.CoverPhoto
					AlternativeSlugs       types.AlternativeSlugs         `json:"alternative_slugs"`
					Breadcrumbs            []types.Breadcrumbs            `json:"breadcrumbs"`
					URLs                   types.URLs                     `json:"urls"`
					Links                  types.Links                    `json:"links"`
					CurrentUserCollections []types.CurrentUserCollections `json:"current_user_collections"`
					TopicSubmissions       types.TopicSubmissions         `json:"topic_submissions"`
					User                   struct {
						types.User
						Links        types.Links        `json:"links"`
						ProfileImage types.ProfileImage `json:"profile_image"`
						Social       types.Social       `json:"social"`
					} `json:"user"`
				} `json:"cover_photo"`
			} `json:"source"`
		} `json:"aggregated"`
	} `json:"tags"`
	Meta types.Meta `json:"meta"`
}

// PublicProfile Get a user’s public profile
// document: https://unsplash.com/documentation#get-a-users-public-profile
func (u *Users) PublicProfile(username string) (reply *UserPublicProfileReply, response *OmitResponse, err error) {
	response, err = u.Client.request("GET", UserGroupReq{}, &reply, nil, username)
	return
}

// PortfolioReply .
type PortfolioReply struct {
	URL string `json:"url"`
}

// Portfolio Get a user’s portfolio link
// document: https://unsplash.com/documentation#get-a-users-portfolio-link
func (u *Users) Portfolio(username string) (reply *PortfolioReply, response *OmitResponse, err error) {
	response, err = u.Client.request("GET", UserGroupReq{}, &reply, nil, username, "portfolio")
	return
}

// UserPhotosReq List a user’s photos request
type UserPhotosReq struct {
	Username    string `url:"username,omitempty"`
	Page        string `url:"page,omitempty"`
	PerPage     string `url:"per_page,omitempty"`
	OrderBy     string `url:"order_by,omitempty"`
	Stats       string `url:"stats,omitempty"`
	Resolution  string `url:"resolution,omitempty"`
	Quantity    string `url:"quantity,omitempty"`
	Orientation string `url:"orientation,omitempty"`
}

// UserPhotosReply List a user’s photos reply
type UserPhotosReply struct {
	types.Photos
	AlternativeSlugs struct {
		types.AlternativeSlugs
	} `json:"alternative_slugs"`
	Breadcrumbs []struct {
		types.Breadcrumbs
	} `json:"breadcrumbs"`
	URLs struct {
		types.URLs
	} `json:"urls"`
	Links struct {
		types.Links
	} `json:"links"`
	CurrentUserCollections []struct {
		types.CurrentUserCollections
	} `json:"current_user_collections"`
	TopicSubmissions struct {
		types.TopicSubmissions
	} `json:"topic_submissions"`
	User struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"user"`
}

func (r UserPhotosReq) API() string {
	return "https://api.unsplash.com/users/"
}

func (r UserPhotosReq) Params() map[string]string {
	return map[string]string{}
}

// Photos List a user’s photos
// Get a list of photos uploaded by a user.
// document: https://unsplash.com/documentation#list-a-users-photos
func (u *Users) Photos(req UserPhotosReq) (reply *[]UserPhotosReply, response *OmitResponse, err error) {
	response, err = u.Client.request("GET", req, &reply, nil, req.Username, "photos")
	return
}

// UserLikesReq List a user’s liked photos request
type UserLikesReq struct {
	Username    string `url:"username,omitempty"`
	Page        string `url:"page,omitempty"`
	PerPage     string `url:"per_page,omitempty"`
	OrderBy     string `url:"order_by,omitempty"`
	Orientation string `url:"orientation,omitempty"`
}

// UserLikesReply List a user’s liked photos reply
type UserLikesReply struct {
	types.Photos
	AlternativeSlugs struct {
		types.AlternativeSlugs
	} `json:"alternative_slugs"`
	Breadcrumbs []struct {
		types.Breadcrumbs
	} `json:"breadcrumbs"`
	URLs struct {
		types.URLs
	} `json:"urls"`
	Links struct {
		types.Links
	} `json:"links"`
	CurrentUserCollections []struct {
		types.CurrentUserCollections
	} `json:"current_user_collections"`
	TopicSubmissions struct {
		types.TopicSubmissions
	} `json:"topic_submissions"`
	User struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"user"`
}

func (UserLikesReq) API() string {
	return "https://api.unsplash.com/users/"
}

func (UserLikesReq) Params() map[string]string {
	return nil
}

// Likes List a user’s liked photos
// Get a list of photos liked by a user.
// document: https://unsplash.com/documentation#list-a-users-liked-photos
func (u *Users) Likes(req UserLikesReq) (reply *[]UserLikesReply, response *OmitResponse, err error) {
	response, err = u.Client.request("GET", req, &reply, nil, req.Username, "likes")
	return
}

// UserCollectionsReq List a user’s collections request
type UserCollectionsReq struct {
	Username string `url:"username,omitempty"`
	Page     string `url:"page,omitempty"`
	PerPage  string `url:"per_page,omitempty"`
}

// UserCollectionsReply List a user’s collections reply
type UserCollectionsReply struct {
	types.Collections
	Tags []struct {
		types.Tags
		Source struct {
			types.Source
			Ancestry struct {
				Type        types.Type        `json:"type"`
				Category    types.Category    `json:"category"`
				Subcategory types.Subcategory `json:"subcategory"`
			} `json:"ancestry"`
			CoverPhoto struct {
				types.CoverPhoto
				AlternativeSlugs       types.AlternativeSlugs         `json:"alternative_slugs"`
				Breadcrumbs            []types.Breadcrumbs            `json:"breadcrumbs"`
				URLs                   types.URLs                     `json:"urls"`
				Links                  types.Links                    `json:"links"`
				CurrentUserCollections []types.CurrentUserCollections `json:"current_user_collections"`
				TopicSubmissions       types.TopicSubmissions         `json:"topic_submissions"`
				User                   struct {
					types.User
					Links        types.Links        `json:"links"`
					ProfileImage types.ProfileImage `json:"profile_image"`
					Social       types.Social       `json:"social"`
				} `json:"user"`
			} `json:"cover_photo"`
		} `json:"source"`
	} `json:"tags"`
	Links types.Links `json:"links"`
	User  struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"user"`
	CoverPhoto struct {
		types.CoverPhoto
		AlternativeSlugs       types.AlternativeSlugs         `json:"alternative_slugs"`
		Breadcrumbs            []types.Breadcrumbs            `json:"breadcrumbs"`
		URLs                   types.URLs                     `json:"urls"`
		Links                  types.Links                    `json:"links"`
		CurrentUserCollections []types.CurrentUserCollections `json:"current_user_collections"`
		TopicSubmissions       types.TopicSubmissions         `json:"topic_submissions"`
		User                   struct {
			types.User
			Links        types.Links        `json:"links"`
			ProfileImage types.ProfileImage `json:"profile_image"`
			Social       types.Social       `json:"social"`
		} `json:"user"`
	} `json:"cover_photo"`

	PreviewPhotos []struct {
		types.PreviewPhotos
		URLs types.URLs `json:"urls"`
	} `json:"preview_photos"`
}

func (UserCollectionsReq) API() string {
	return "https://api.unsplash.com/users/"
}

func (UserCollectionsReq) Params() map[string]string {
	return nil
}

// Collections List a user’s collections
// Get a list of collections created by the user.
// documents: https://unsplash.com/documentation#list-a-users-collections
func (u *Users) Collections(req UserCollectionsReq) (reply *[]UserCollectionsReply, response *OmitResponse, err error) {
	response, err = u.Client.request("GET", req, &reply, nil, req.Username, "collections")
	return
}

// UserStatisticsReq Get a user’s statistics request
type UserStatisticsReq struct {
	Username   string `url:"username,omitempty"`
	Resolution string `url:"resolution,omitempty"`
	Quantity   string `url:"quantity,omitempty"`
}

// UserStatisticsReply Get a user’s statistics reply
type UserStatisticsReply struct {
	types.Statistics
	Downloads struct {
		types.Downloads
		Historical struct {
			types.Historical
			Values []types.Values `json:"values"`
		} `json:"historical"`
	} `json:"downloads"`
	Views struct {
		types.Views
		Historical struct {
			types.Historical
			Values []types.Values `json:"values"`
		} `json:"historical"`
	} `json:"views"`
}

func (UserStatisticsReq) API() string {
	return "https://api.unsplash.com/users/"
}

func (UserStatisticsReq) Params() map[string]string {
	return nil
}

// Statistics Get a user’s statistics
// Retrieve the consolidated number of downloads, views and likes of all user’s photos,
// as well as the historical breakdown and average of these stats in a specific timeframe (default is 30 days).
// document: https://unsplash.com/documentation#get-a-users-statistics
func (u *Users) Statistics(req UserStatisticsReq) (reply *UserStatisticsReply, response *OmitResponse, err error) {
	response, err = u.Client.request("GET", req, &reply, nil, req.Username, "statistics")
	return
}
