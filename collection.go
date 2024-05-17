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

type Collection struct {
	Client *Client
}

// CollectionReq List collections request
type CollectionReq struct {
	Page    string `url:"page,omitempty"`
	PerPage string `url:"per_page,omitempty"`
}

// CollectionReply List collections reply
type CollectionReply struct {
	types.Collections
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

func (CollectionReq) API() string {
	return "https://api.unsplash.com/collections"
}

func (CollectionReq) Params() map[string]string {
	return nil
}

// Collections List collections
// Get a single page from the list of all collections.
// document: https://unsplash.com/documentation#list-collections
func (c *Collection) Collections(req CollectionReq) (reply []*CollectionReply, response *OmitResponse, err error) {
	response, err = c.Client.request("GET", req, &reply, nil)
	return
}

// GetACollectionReq Get a collection request
type GetACollectionReq struct {
}

// GetACollectionReply Get a collection reply
type GetACollectionReply struct {
	types.Collections
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
	Meta types.Meta `json:"meta"`
}

func (GetACollectionReq) API() string {
	return "https://api.unsplash.com/collections"
}

func (GetACollectionReq) Params() map[string]string {
	return nil
}

// GetACollection Get a collection
// Retrieve a single collection. To view a user’s private collections, the read_collections scope is required.
// document: https://unsplash.com/documentation#get-a-collections-photos
func (c *Collection) GetACollection(collectionId string) (reply *GetACollectionReply, response *OmitResponse, err error) {
	response, err = c.Client.request("GET", GetACollectionReq{}, &reply, nil, collectionId)
	return
}

// GetACollectionPhotosReq Get a collection’s photos request
type GetACollectionPhotosReq struct {
	ID          string `url:"id"`
	Page        string `url:"page,omitempty"`
	PerPage     string `url:"per_page,omitempty"`
	Orientation string `url:"orientation,omitempty"`
}

// GetACollectionPhotosReply Get a collection’s photos reply
type GetACollectionPhotosReply struct {
	types.Photos
	AlternativeSlugs       types.AlternativeSlugs         `json:"alternative_slugs"`
	Breadcrumbs            []types.Breadcrumbs            `json:"breadcrumbs"`
	URLs                   types.URLs                     `json:"urls"`
	Links                  types.Links                    `json:"links"`
	CurrentUserCollections []types.CurrentUserCollections `json:"current_user_collections"`
	TopicSubmissions       struct {
		types.TopicSubmissions
	} `json:"topic_submissions"`
	User struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"user"`
}

func (GetACollectionPhotosReq) API() string {
	return "https://api.unsplash.com/collections"
}

func (GetACollectionPhotosReq) Params() map[string]string {
	return nil
}

// GetACollectionPhotos Get a collection’s photos
// document: https://unsplash.com/documentation#get-a-collections-photos
func (c *Collection) GetACollectionPhotos(req GetACollectionPhotosReq) (reply []*GetACollectionPhotosReply, response *OmitResponse, err error) {
	response, err = c.Client.request("GET", GetACollectionReq{}, &reply, nil, req.ID, "photos")
	return
}

// RelatedCollectionsReq List a collection’s related collections request
type RelatedCollectionsReq struct {
}

// RelatedCollectionsReply List a collection’s related collections reply
type RelatedCollectionsReply struct {
	types.Collections
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

func (RelatedCollectionsReq) API() string {
	return "https://api.unsplash.com/collections"
}

func (RelatedCollectionsReq) Params() map[string]string {
	return nil
}

// RelatedCollections List a collection’s related collections
// Retrieve a list of collections related to this one.
// document: https://unsplash.com/documentation#list-a-collections-related-collections
func (c *Collection) RelatedCollections(collectionId string) (reply []*RelatedCollectionsReply, response *OmitResponse, err error) {
	response, err = c.Client.request("GET", RelatedCollectionsReq{}, &reply, nil, collectionId, "related")
	return
}

// NewCollectionsReq Create a new collection request
type NewCollectionsReq struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Private     bool   `json:"private,omitempty"`
}

// NewCollectionsReply Create a new collection reply
type NewCollectionsReply struct {
	types.Collections
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
}

func (NewCollectionsReq) API() string {
	return "https://api.unsplash.com/collections"
}

func (NewCollectionsReq) Params() map[string]string {
	return nil
}

// NewCollections Create a new collection
// Create a new collection. This requires the write_collections scope.
// Note: you must be activate your account first
// document: https://unsplash.com/documentation#create-a-new-collection
func (c *Collection) NewCollections(token string, req NewCollectionsReq) (reply NewCollectionsReply, response *OmitResponse, err error) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	response, err = c.Client.request("POST", req, &reply, headers)
	return
}

// UpdateCollectionReq Update an existing collection request
type UpdateCollectionReq struct {
	ID          string `json:"id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Private     bool   `json:"private,omitempty"`
}

// UpdateCollectionReply Update an existing collection reply
type UpdateCollectionReply struct {
	types.Collections
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
}

func (UpdateCollectionReq) API() string {
	return "https://api.unsplash.com/collections"
}

func (UpdateCollectionReq) Params() map[string]string {
	return nil
}

// UpdateCollection Update an existing collection
// Update an existing collection belonging to the logged-in user. This requires the `write_collections` scope.
// document: https://unsplash.com/documentation#update-an-existing-collection
func (c *Collection) UpdateCollection(token string, req UpdateCollectionReq) (reply *UpdateCollectionReply, response *OmitResponse, err error) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	response, err = c.Client.request("PUT", req, &reply, headers, req.ID)
	return
}

// DeleteCollectionReq Delete a collection request
type DeleteCollectionReq struct {
}

// DeleteCollectionReply Delete a collection reply
type DeleteCollectionReply struct {
}

func (DeleteCollectionReq) API() string {
	return "https://api.unsplash.com/collections"
}

func (DeleteCollectionReq) Params() map[string]string {
	return nil
}

// DeleteCollection Delete a collection
// Delete a collection belonging to the logged-in user. This requires the `write_collections` scope.
// document: https://unsplash.com/documentation#delete-a-collection
func (c *Collection) DeleteCollection(token string, collectionId string) (response *OmitResponse, err error) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	response, err = c.Client.request("DELETE", DeleteCollectionReq{}, nil, headers, collectionId)
	return
}

// AddPhotoReq Add a photo to a collection request
type AddPhotoReq struct {
	CollectionId string `json:"collection_id"`
	PhotoId      string `json:"photo_id"`
}

// AddPhotoReply Add a photo to a collection reply
type AddPhotoReply struct {
	CreatedAt string `json:"created_at"`
	Photo     struct {
		types.Photos
		AlternativeSlugs       types.AlternativeSlugs         `json:"alternative_slugs"`
		Breadcrumbs            []types.Breadcrumbs            `json:"breadcrumbs"`
		URLs                   types.URLs                     `json:"urls"`
		Links                  types.Links                    `json:"links"`
		CurrentUserCollections []types.CurrentUserCollections `json:"current_user_collections"`
		TopicSubmissions       struct {
			types.TopicSubmissions
		} `json:"topic_submissions"`
		User struct {
			types.User
			Links        types.Links        `json:"links"`
			ProfileImage types.ProfileImage `json:"profile_image"`
			Social       types.Social       `json:"social"`
		} `json:"user"`
	} `json:"photo"`
	Collection struct {
		types.Collections
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
	} `json:"collection"`
	User struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"user"`
}

func (AddPhotoReq) API() string {
	return "https://api.unsplash.com/collections"
}

func (AddPhotoReq) Params() map[string]string {
	return nil
}

// AddPhoto Add a photo to a collection
// Add a photo to one of the logged-in user’s collections. Requires the `write_collections` scope.
// document: https://unsplash.com/documentation#add-a-photo-to-a-collection
func (c *Collection) AddPhoto(token string, req AddPhotoReq) (reply *AddPhotoReply, response *OmitResponse, err error) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	response, err = c.Client.request("POST", req, &reply, headers, req.CollectionId, "add")
	return
}

// RemovePhotoReq Remove a photo from a collection request
type RemovePhotoReq struct {
	CollectionId string `json:"collection_id"`
	PhotoId      string `json:"photo_id"`
}

// RemovePhotoReply Remove a photo from a collection reply
type RemovePhotoReply struct {
	CreatedAt string `json:"created_at"`
	Photo     struct {
		types.Photos
		AlternativeSlugs       types.AlternativeSlugs         `json:"alternative_slugs"`
		Breadcrumbs            []types.Breadcrumbs            `json:"breadcrumbs"`
		URLs                   types.URLs                     `json:"urls"`
		Links                  types.Links                    `json:"links"`
		CurrentUserCollections []types.CurrentUserCollections `json:"current_user_collections"`
		TopicSubmissions       struct {
			types.TopicSubmissions
		} `json:"topic_submissions"`
		User struct {
			types.User
			Links        types.Links        `json:"links"`
			ProfileImage types.ProfileImage `json:"profile_image"`
			Social       types.Social       `json:"social"`
		} `json:"user"`
	} `json:"photo"`
	Collection struct {
		types.Collections
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
	} `json:"collection"`
	User struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"user"`
}

func (RemovePhotoReq) API() string {
	return "https://api.unsplash.com/collections"
}

func (RemovePhotoReq) Params() map[string]string {
	return nil
}

// RemovePhoto Remove a photo from a collection
// Remove a photo from one of the logged-in user’s collections. Requires the `write_collections` scope.
// document: https://unsplash.com/documentation#remove-a-photo-from-a-collection
func (c *Collection) RemovePhoto(token string, req RemovePhotoReq) (reply *RemovePhotoReply, response *OmitResponse, err error) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	response, err = c.Client.request("DELETE", req, &reply, headers, req.CollectionId, "remove")
	return
}
