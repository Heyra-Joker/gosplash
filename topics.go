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

type Topics struct {
	Client *Client
}

// TopicsListReq List topics request
type TopicsListReq struct {
	IDs     string `url:"i_ds,omitempty"`
	Page    string `url:"page,omitempty"`
	PerPage string `url:"per_page,omitempty"`
	OrderBy string `url:"order_by,omitempty"`
}

// TopicsListReply List topics reply
type TopicsListReply struct {
	types.Topics
	CurrentUserContributions []types.CurrentUserContributions `json:"current_user_contributions"`
	Links                    types.Links                      `json:"links"`
	Owners                   []struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"owners"`
	CoverPhoto struct {
		types.CoverPhoto
		AlternativeSlugs       types.AlternativeSlugs         `json:"alternative_slugs"`
		Breadcrumbs            []types.Breadcrumbs            `json:"breadcrumbs"`
		URLs                   types.URLs                     `json:"urls"`
		Links                  types.Links                    `json:"links"`
		CurrentUserCollections []types.CurrentUserCollections `json:"current_user_collections"`
		TopicSubmissions       struct {
			types.TopicSubmissions
			UGC       types.UGC       `json:"ugc"`
			Travel    types.Travel    `json:"travel"`
			Nature    types.Nature    `json:"nature"`
			CoolTones types.CoolTones `json:"cool_tones"`
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

func (TopicsListReq) API() string {
	return "https://api.unsplash.com/topics"
}

func (TopicsListReq) Params() map[string]string {
	return nil
}

// TopicsList List topics
// Get a single page from the list of all topics.
// document: https://unsplash.com/documentation#list-topics
func (t *Topics) TopicsList(req TopicsListReq) (reply []*TopicsListReply, response *OmitResponse, err error) {
	response, err = t.Client.request("GET", req, &reply, nil)
	return
}

// GetATopicReq Get a topic request
type GetATopicReq struct {
}

// GetATopicReply Get a topic reply
type GetATopicReply struct {
	types.Topics
	CurrentUserContributions []types.CurrentUserContributions `json:"current_user_contributions"`
	Links                    types.Links                      `json:"links"`
	Owners                   []struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"owners"`
	TopContributors []struct {
		types.TopContributors
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"top_contributors"`
	CoverPhoto struct {
		types.CoverPhoto
		AlternativeSlugs       types.AlternativeSlugs         `json:"alternative_slugs"`
		Breadcrumbs            []types.Breadcrumbs            `json:"breadcrumbs"`
		URLs                   types.URLs                     `json:"urls"`
		Links                  types.Links                    `json:"links"`
		CurrentUserCollections []types.CurrentUserCollections `json:"current_user_collections"`
		TopicSubmissions       struct {
			types.TopicSubmissions
			UGC       types.UGC       `json:"ugc"`
			Travel    types.Travel    `json:"travel"`
			Nature    types.Nature    `json:"nature"`
			CoolTones types.CoolTones `json:"cool_tones"`
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

func (GetATopicReq) API() string {
	return "https://api.unsplash.com/topics"
}

func (GetATopicReq) Params() map[string]string {
	return nil
}

// GetATopic Get a topic
// Retrieve a single topic.
// document: https://unsplash.com/documentation#get-a-topic
func (t *Topics) GetATopic(idOrSlug string) (reply *GetATopicReply, response *OmitResponse, err error) {
	response, err = t.Client.request("GET", GetATopicReq{}, &reply, nil, idOrSlug)
	return
}

// TopicPhotosReq Get a topic’s photos request
type TopicPhotosReq struct {
	IdOrSlug    string `url:"id_or_slug"`
	Page        string `url:"page,omitempty"`
	PerPage     string `url:"per_page,omitempty"`
	Orientation string `url:"orientation,omitempty"`
	OrderBy     string `url:"order_by,omitempty"`
}

// TopicPhotosReply Get a topic’s photos reply
type TopicPhotosReply struct {
	types.Photos
	AlternativeSlugs       types.AlternativeSlugs `json:"alternative_slugs"`
	Breadcrumbs            []types.Breadcrumbs    `json:"breadcrumbs"`
	URLs                   types.URLs             `json:"urls"`
	Links                  types.Links            `json:"links"`
	CurrentUserCollections []struct {
		types.CurrentUserCollections
	} `json:"current_user_collections"`
	TopicSubmissions struct {
		types.TopicSubmissions
		UGC       types.UGC       `json:"ugc"`
		Travel    types.Travel    `json:"travel"`
		Nature    types.Nature    `json:"nature"`
		CoolTones types.CoolTones `json:"cool_tones"`
	} `json:"topic_submissions"`
	User struct {
		types.User
		Links        types.Links        `json:"links"`
		ProfileImage types.ProfileImage `json:"profile_image"`
		Social       types.Social       `json:"social"`
	} `json:"user"`
}

func (TopicPhotosReq) API() string {
	return "https://api.unsplash.com/topics"
}

func (TopicPhotosReq) Params() map[string]string {
	return nil
}

// TopicPhotos Get a topic’s photos
// Retrieve a topic’s photos.
// document: https://unsplash.com/documentation#get-a-topics-photos
func (t *Topics) TopicPhotos(req TopicPhotosReq) (reply []*TopicPhotosReply, response *OmitResponse, err error) {
	response, err = t.Client.request("GET", req, &reply, nil, req.IdOrSlug, "photos")
	return
}
