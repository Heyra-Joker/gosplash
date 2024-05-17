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

type Stats struct {
	TotalPhotos        int64 `json:"total_photos"`
	PhotoDownloads     int64 `json:"photo_downloads"`
	Photos             int64 `json:"photos"`
	Downloads          int64 `json:"downloads"`
	Views              int64 `json:"views"`
	Photographers      int64 `json:"photographers"`
	Pixels             int64 `json:"pixels"`
	DownloadsPerSecond int64 `json:"downloads_per_second"`
	ViewsPerSecond     int64 `json:"views_per_second"`
	Developers         int64 `json:"developers"`
	Applications       int64 `json:"applications"`
	Requests           int64 `json:"requests"`
}
