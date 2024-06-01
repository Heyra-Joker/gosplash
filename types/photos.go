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

package types

type Photos struct {
	ID             string `json:"id"`
	Slug           string `json:"slug"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	BlurHash       string `json:"blur_hash"`
	AssetType      string `json:"asset_type"`
	PromotedAt     string `json:"promoted_at"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Color          string `json:"color"`
	Description    string `json:"description"`
	AltDescription string `json:"alt_description"`
	Likes          int    `json:"likes"`
	LikedByUser    bool   `json:"liked_by_user"`
	//Sponsorship    string `json:"sponsorship"`
	PublicDomain bool `json:"public_domain"`
	Views        int  `json:"views"`
	Downloads    int  `json:"downloads"`
}
