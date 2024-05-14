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

type Links struct {
	Self             string `json:"self"`
	HTML             string `json:"html"`
	Photos           string `json:"photos"`
	Likes            string `json:"likes"`
	Portfolio        string `json:"portfolio"`
	Following        string `json:"following"`
	Followers        string `json:"followers"`
	Download         string `json:"download"`
	DownloadLocation string `json:"download_location"`
}
