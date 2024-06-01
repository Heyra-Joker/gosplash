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

type Results struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	PublishedAt     string `json:"published_at"`
	LastCollectedAt string `json:"last_collected_at"`
	UpdatedAt       string `json:"updated_at"`
	Featured        bool   `json:"featured"`
	TotalPhotos     int    `json:"total_photos"`
	Private         bool   `json:"private"`
	ShareKey        string `json:"share_key"`
	Slug            string `json:"slug"`
	CreatedAt       string `json:"created_at"`
	PromotedAt      string `json:"promoted_at"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	Color           string `json:"color"`
	BlurHash        string `json:"blur_hash"`
	AltDescription  string `json:"alt_description"`
	Likes           int    `json:"likes"`
	LikedByUser     bool   `json:"liked_by_user"`
	//Sponsorship     string `json:"sponsorship"`
	AssetType string `json:"asset_type"`
}
