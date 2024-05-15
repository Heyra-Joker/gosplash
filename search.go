/*
 * Copyright @2024
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gosplash

import "github.com/Heyra-Joker/gosplash/types"

// Search Group
type Search struct {
	Client *Client
}

// SearchPhotosReq Search photos request
type SearchPhotosReq struct {
	Query         string `url:"query"`
	Page          string `url:"page,omitempty"`
	PerPage       string `url:"per_page,omitempty"`
	OrderBy       string `url:"order_by,omitempty"`
	Collections   string `url:"collections,omitempty"`
	ContentFilter string `url:"content_filter,omitempty"`
	Color         string `url:"color,omitempty"`
	Orientation   string `url:"orientation,omitempty"`
	// Beta parameters (for access to beta parameters, email api@unsplash.com with your application ID
	Lang string `url:"lang,omitempty"`
}

// SearchPhotosReply Search photos reply
type SearchPhotosReply struct {
	types.Search
	Results []struct {
		types.Results
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
		Tags []struct {
			types.Tags
			Source struct {
				types.Source
				Ancestry struct {
					types.Ancestry
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
	} `json:"results"`
}

func (SearchPhotosReq) API() string {
	return "https://api.unsplash.com/search/photos"
}

func (SearchPhotosReq) Params() map[string]string {
	return nil
}

// Photos Search photos
// Get a single page of photo results for a query.
// Note: Beta parameters (for access to beta parameters, email api@unsplash.com with your application ID):
// document: https://unsplash.com/documentation#search-photos
func (s *Search) Photos(req SearchPhotosReq) (reply *SearchPhotosReply, response *OmitResponse, err error) {
	response, err = s.Client.request("GET", req, &reply, nil)
	return
}

// SearchCollectionsReq Search collections request
type SearchCollectionsReq struct {
	Query   string `url:"query"`
	Page    string `url:"page,omitempty"`
	PerPage string `url:"per_page,omitempty"`
}

// SearchCollectionsReply Search collections reply
type SearchCollectionsReply struct {
	types.Search
	Results []struct {
		types.Results
		Tags []struct {
			types.Tags
			Source struct {
				types.Source
				Ancestry struct {
					types.Ancestry
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
			TopicSubmissions       struct {
				types.TopicSubmissions
				Wallpapers   types.Wallpapers   `json:"wallpapers"`
				Animals      types.Animals      `json:"animals"`
				Blue         types.Blue         `json:"blue"`
				Experimental types.Experimental `json:"experimental"`
			} `json:"topic_submissions"`
			User struct {
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
}

func (SearchCollectionsReq) API() string {
	return "https://api.unsplash.com/search/collections"
}

func (SearchCollectionsReq) Params() map[string]string {
	return nil
}

// Collections Search collections
// Get a single page of collection results for a query.
// document: https://unsplash.com/documentation#search-collections
func (s *Search) Collections(req SearchCollectionsReq) (reply *SearchCollectionsReply, response *OmitResponse, err error) {
	response, err = s.Client.request("GET", req, &reply, nil)
	return
}

// SearchUsersReq Search users request
type SearchUsersReq struct {
	Query   string `url:"query"`
	Page    string `url:"page,omitempty"`
	PerPage string `url:"per_page,omitempty"`
}

// SearchUsersReply Search users reply
type SearchUsersReply struct {
	types.Search
	Results []struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
		Photos       []struct {
			types.Photos
			URLs types.URLs `json:"urls"`
		} `json:"photos"`
	} `json:"results"`
}

func (SearchUsersReq) API() string {
	return "https://api.unsplash.com/search/users"
}

func (SearchUsersReq) Params() map[string]string {
	return nil
}

// Users Search users
// Get a single page of user results for a query.
// document: https://unsplash.com/documentation#search-users
func (s *Search) Users(req SearchUsersReq) (reply *SearchUsersReply, response *OmitResponse, err error) {
	response, err = s.Client.request("GET", req, &reply, nil)
	return
}
