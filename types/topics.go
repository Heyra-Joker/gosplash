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

type Topics struct {
	Id                          string `json:"id"`
	Slug                        string `json:"slug"`
	Title                       string `json:"title"`
	Description                 string `json:"description"`
	PublishedAt                 string `json:"published_at"`
	UpdatedAt                   string `json:"updated_at"`
	StartsAt                    string `json:"starts_at"`
	EndsAt                      string `json:"ends_at"`
	OnlySubmissionsAfter        string `json:"only_submissions_after"`
	Visibility                  string `json:"visibility"`
	Featured                    bool   `json:"featured"`
	TotalPhotos                 int    `json:"total_photos"`
	TotalCurrentUserSubmissions string `json:"total_current_user_submissions"`
	Status                      string `json:"status"`
}
