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

type TopContributors struct {
	Id                         string `json:"id"`
	UpdatedAt                  string `json:"updated_at"`
	Username                   string `json:"username"`
	Name                       string `json:"name"`
	FirstName                  string `json:"first_name"`
	LastName                   string `json:"last_name"`
	TwitterUsername            string `json:"twitter_username"`
	PortfolioUrl               string `json:"portfolio_url"`
	Bio                        string `json:"bio"`
	Location                   string `json:"location"`
	InstagramUsername          string `json:"instagram_username"`
	TotalCollections           int    `json:"total_collections"`
	TotalLikes                 int    `json:"total_likes"`
	TotalPhotos                int    `json:"total_photos"`
	TotalPromotedPhotos        int    `json:"total_promoted_photos"`
	TotalIllustrations         int    `json:"total_illustrations"`
	TotalPromotedIllustrations int    `json:"total_promoted_illustrations"`
	AcceptedTos                bool   `json:"accepted_tos"`
	ForHire                    bool   `json:"for_hire"`
}
