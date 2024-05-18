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

import (
	"fmt"
	"github.com/Heyra-Joker/gosplash/types"
)

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
func (p *Photo) Random(req PhotoRandomReq) (reply []*PhotoRandomReply, response *OmitResponse, err error) {
	if req.Count == "" {
		req.Count = "1"
	}
	response, err = p.Client.request("GET", req, &reply, nil)
	return
}

// PhotoStatisticsReq Get a photo’s statistics request
type PhotoStatisticsReq struct {
}

// PhotoStatisticsReply Get a photo’s statistics reply
type PhotoStatisticsReply struct {
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
	Likes struct {
		types.Likes
		Historical struct {
			types.Historical
			Values []types.Values `json:"values"`
		} `json:"historical"`
	} `json:"likes"`
}

func (PhotoStatisticsReq) API() string {
	return "https://api.unsplash.com/photos"
}

func (PhotoStatisticsReq) Params() map[string]string {
	return nil
}

// Statistics Get a photo’s statistics
// Retrieve total number of downloads, views and likes of a single photo,
// as well as the historical breakdown of these stats in a specific timeframe (default is 30 days).
// document: https://unsplash.com/documentation#get-a-photos-statistics
func (p *Photo) Statistics(photoId string) (reply *PhotoStatisticsReply, response *OmitResponse, err error) {
	response, err = p.Client.request("GET", PhotoStatisticsReq{}, &reply, nil, photoId, "statistics")
	return
}

// PhotoTrackDownloadReq Track a photo download request
type PhotoTrackDownloadReq struct {
}

// PhotoTrackDownloadReply Track a photo download reply
type PhotoTrackDownloadReply struct {
	URL string `json:"url"`
}

func (PhotoTrackDownloadReq) API() string {
	return "https://api.unsplash.com/photos"
}

func (PhotoTrackDownloadReq) Params() map[string]string {
	return nil
}

// TrackDownload Track a photo download
// Note: some photo can not download, error will return `{"errors":["Unauthorized"]}`
// document: https://unsplash.com/documentation#track-a-photo-download
func (p *Photo) TrackDownload(photoId string) (reply *PhotoTrackDownloadReply, response *OmitResponse, err error) {
	response, err = p.Client.request("GET", PhotoTrackDownloadReq{}, &reply, nil, photoId, "download")
	return
}

// PhotoUpdateReq Update a photo on behalf of the logged-in user. request
type PhotoUpdateReq struct {
	ID                  string `json:"id"`
	Description         string `json:"description,omitempty"`
	ShowOnProfile       string `json:"show_on_profile,omitempty"`
	Tags                string `json:"tags,omitempty"`
	LocationLatitude    string `json:"location[latitude],omitempty"`
	LocationLongitude   string `json:"location[longitude],omitempty"`
	LocationName        string `json:"location[name],omitempty"`
	LocationCity        string `json:"location[city],omitempty"`
	LocationCountry     string `json:"location[country],omitempty"`
	ExifMake            string `json:"exif[make],omitempty"`
	ExifModel           string `json:"exif[model],omitempty"`
	ExifExposureTime    string `json:"exif[exposure_time],omitempty"`
	ExifApertureValue   string `json:"exif[aperture_value],omitempty"`
	ExifFocalLength     string `json:"exif[focal_length],omitempty"`
	ExifIsoSpeedRatings string `json:"exif[iso_speed_ratings],omitempty"`
}

// PhotoUpdateReply Update a photo on behalf of the logged-in user reply
type PhotoUpdateReply struct {
	types.Photos
	Location struct {
		types.Location
		Position types.Position `json:"position"`
	} `json:"location"`
	Tags                   []types.Tags `json:"tags"`
	CurrentUserCollections []struct {
		types.CurrentUserCollections
		CoverPhoto types.CoverPhoto `json:"cover_photo"`
		User       types.User       `json:"user"`
	} `json:"current_user_collections"`
	URLs  types.URLs  `json:"urls"`
	Links types.Links `json:"links"`
	User  struct {
		types.User
		Links types.Links `json:"links"`
	} `json:"user"`
}

func (PhotoUpdateReq) API() string {
	return "https://api.unsplash.com/photos"
}

func (PhotoUpdateReq) Params() map[string]string {
	return nil
}

// UnimplementedUpdate a photo on behalf of the logged-in user. This requires the write_photos scope.
// Note: Sorry about this function, I have no permission to upload image :(
// If you have the necessary permissions, you can submit an issue to fix this function.
// document: https://unsplash.com/documentation#update-a-photo
func (p *Photo) UnimplementedUpdate(token string, req PhotoUpdateReq) (reply *PhotoUpdateReply, response *OmitResponse, err error) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	response, err = p.Client.request("PUT", req, &reply, headers, req.ID)
	return
}

// PhotoLikeReq Like a photo request
type PhotoLikeReq struct {
}

// PhotoLikeReply Like a photo reply
type PhotoLikeReply struct {
	Photo struct {
		types.Photos
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
	} `json:"photo"`
	User struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"user"`
}

func (PhotoLikeReq) API() string {
	return "https://api.unsplash.com/photos"
}

func (PhotoLikeReq) Params() map[string]string {
	return nil
}

// Like  a photo
// Like a photo on behalf of the logged-in user. This requires the write_likes scope.
// document: https://unsplash.com/documentation#like-a-photo
func (p *Photo) Like(token string, photoId string) (reply *PhotoLikeReply, response *OmitResponse, err error) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	response, err = p.Client.request("POST", PhotoLikeReq{}, &reply, headers, photoId, "like")
	return
}

// PhotoUnLikeReq unlike a photo request
type PhotoUnLikeReq struct {
}

// PhotoUnLikeReply unlike a photo reply
type PhotoUnLikeReply struct {
	Photo struct {
		types.Photos
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
	} `json:"photo"`
	User struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"user"`
}

func (PhotoUnLikeReq) API() string {
	return "https://api.unsplash.com/photos"
}

func (PhotoUnLikeReq) Params() map[string]string {
	return nil
}

// UnLike  a photo
// UnLike a photo on behalf of the logged-in user. This requires the write_likes scope.
// document: https://unsplash.com/documentation#like-a-photo
func (p *Photo) UnLike(token string, photoId string) (reply *PhotoUnLikeReply, response *OmitResponse, err error) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	response, err = p.Client.request("DELETE", PhotoUnLikeReq{}, &reply, headers, photoId, "like")
	return
}
