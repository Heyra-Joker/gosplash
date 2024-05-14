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

import (
	"fmt"
	"github.com/Heyra-Joker/gosplash/types"
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
	types.User
	Links        types.Links        `json:"links"`
	ProfileImage types.ProfileImage `json:"profile_image"`
	Social       types.Social       `json:"social"`
	Tags         types.Tags         `json:"tags"`
	Meta         types.Meta         `json:"meta"`
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
func (c *CurrentUser) Me(token string) (reply *MeReply, response *OmitResponse, err error) {
	reqHeaders := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	response, err = c.Client.request("GET", MeRequest{}, &reply, reqHeaders)
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
	types.User
	Links        types.Links        `json:"links"`
	ProfileImage types.ProfileImage `json:"profile_image"`
	Social       types.Social       `json:"social"`
	Tags         types.Tags         `json:"tags"`
	Meta         types.Meta         `json:"meta"`
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
func (c *CurrentUser) UpdateMe(token string, req UpdateMeRequest) (reply *UpdateMeReply, headers *OmitResponse, err error) {
	reqHeaders := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	headers, err = c.Client.request("PUT", req, &reply, reqHeaders)
	return
}
