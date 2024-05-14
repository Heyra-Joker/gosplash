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

package main

import "github.com/Heyra-Joker/gosplash/types"

type Photo struct {
	Client *Client
}

// PhotosReq List photos request
type PhotosReq struct {
	Page    string `url:"page,omitempty"`
	PerPage string `url:"per_page,omitempty"`
	OrderBy string `url:"order_by,omitempty"`
}

// PhotosReply List photos reply
type PhotosReply struct {
	types.Photos
	AlternativeSlugs types.AlternativeSlugs `json:"alternative_slugs"`
	Breadcrumbs      []struct {
		types.Breadcrumbs
	} `json:"breadcrumbs"`
	URLs                   types.URLs  `json:"urls"`
	Links                  types.Links `json:"links"`
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

func (PhotosReq) API() string {
	return "https://api.unsplash.com/photos"
}

func (PhotosReq) Params() map[string]string {
	return nil
}

// Photos List photos
// Get a single page from the Editorial feed.
// document: https://unsplash.com/documentation#list-photos
func (p *Photo) Photos(req PhotosReq) (reply []*PhotosReply, response *OmitResponse, err error) {
	response, err = p.Client.request("GET", req, &reply, nil)
	return
}

// GetAPhotoReq Get a photo request
type GetAPhotoReq struct {
}

// GetAPhotoReply Get a photo reply
type GetAPhotoReply struct {
	types.Photos
	AlternativeSlugs types.AlternativeSlugs `json:"alternative_slugs"`
	Breadcrumbs      []struct {
		types.Breadcrumbs
	} `json:"breadcrumbs"`
	URLs                   types.URLs  `json:"urls"`
	Links                  types.Links `json:"links"`
	CurrentUserCollections []struct {
		types.CurrentUserCollections
	} `json:"current_user_collections"`
	TopicSubmissions types.TopicSubmissions `json:"topic_submissions"`
	User             struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"user"`
	Exif     types.Exif `json:"exif"`
	Location struct {
		types.Location
		Position types.Position `json:"position"`
	} `json:"location"`
	Meta types.Meta `json:"meta"`
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

	TagsPreview []struct {
		types.TagsPreview
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
	} `json:"tags_preview"`

	Topics []types.Topics `json:"topics"`

	RelatedCollections struct {
		types.RelatedCollections
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
		} `json:"results"`
	} `json:"related_collections"`
}

func (GetAPhotoReq) API() string {
	return "https://api.unsplash.com/photos"
}

func (GetAPhotoReq) Params() map[string]string {
	return nil
}

// GetAPhoto Get a photo
// Retrieve a single photo.
// document: https://unsplash.com/documentation#get-a-photo
func (p *Photo) GetAPhoto(photoId string) (reply *GetAPhotoReply, response *OmitResponse, err error) {
	response, err = p.Client.request("GET", GetAPhotoReq{}, &reply, nil, photoId)
	return
}

// PhotoRandomReq Get a random photo request
type PhotoRandomReq struct {
	Collections   string `url:"collections,omitempty"`
	Topics        string `url:"topics,omitempty"`
	Username      string `url:"username,omitempty"`
	Query         string `url:"query,omitempty"`
	Orientation   string `url:"orientation,omitempty"`
	ContentFilter string `url:"content_filter,omitempty"`
	Count         string `url:"count,omitempty"`
}

// PhotoRandomReply Get a random photo reply
type PhotoRandomReply struct {
	types.Photos
	AlternativeSlugs types.AlternativeSlugs `json:"alternative_slugs"`
	Breadcrumbs      []struct {
		types.Breadcrumbs
	} `json:"breadcrumbs"`
	URLs                   types.URLs  `json:"urls"`
	Links                  types.Links `json:"links"`
	CurrentUserCollections []struct {
		types.CurrentUserCollections
	} `json:"current_user_collections"`
	TopicSubmissions types.TopicSubmissions `json:"topic_submissions"`
	User             struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"user"`
	Exif     types.Exif `json:"exif"`
	Location struct {
		types.Location
		Position types.Position `json:"position"`
	} `json:"location"`
	Meta types.Meta `json:"meta"`
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
	TagsPreview []struct {
		types.TagsPreview
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
	} `json:"tags_preview"`
	Topics []types.Topics `json:"topics"`
}

func (PhotoRandomReq) API() string {
	return "https://api.unsplash.com/photos/random"
}

func (PhotoRandomReq) Params() map[string]string {
	return nil
}

// Random Get a random photo
// Retrieve a single random photo, given optional filters.
// All parameters are optional, and can be combined to narrow the pool of photos from which a random one will be chosen.
// document: https://unsplash.com/documentation#get-a-random-photo
func (p *Photo) Random(req PhotoRandomReq) (reply *PhotoRandomReply, response *OmitResponse, err error) {
	response, err = p.Client.request("GET", req, &reply, nil)
	return
}
