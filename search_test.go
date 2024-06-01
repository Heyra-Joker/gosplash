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
	"encoding/json"
	"testing"
)

func TestSearch_Photos(t *testing.T) {
	searchGroup := Search{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := searchGroup.Photos(SearchPhotosReq{
		Query:         "beauty women",
		Page:          "1",
		PerPage:       "1",
		OrderBy:       "",
		Collections:   "",
		ContentFilter: "high",
		Color:         "",
		Orientation:   "",
		Lang:          "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}

func TestSearch_Collections(t *testing.T) {
	searchGroup := Search{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := searchGroup.Collections(SearchCollectionsReq{
		Query:   "cat",
		Page:    "1",
		PerPage: "1",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}

func TestSearch_Users(t *testing.T) {
	searchGroup := Search{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := searchGroup.Users(SearchUsersReq{
		Query:   "thanks",
		Page:    "1",
		PerPage: "1",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}
