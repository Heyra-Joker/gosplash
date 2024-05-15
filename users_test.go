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

func TestUsers_PublicProfile(t *testing.T) {
	userGroup := Users{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := userGroup.PublicProfile("br00m")
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(res)
	t.Log(string(data))
	t.Log(response.OriginResponseBody)
}

func TestUsers_Portfolio(t *testing.T) {
	userGroup := Users{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, responseHeaders, err := userGroup.Portfolio("br00m")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", res)
	t.Logf("%#v", responseHeaders)

}

func TestUsers_Photos(t *testing.T) {
	userGroup := Users{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := userGroup.Photos(UserPhotosReq{
		Username:    "br00m",
		Page:        "1",
		PerPage:     "1",
		OrderBy:     "",
		Stats:       "",
		Resolution:  "",
		Quantity:    "",
		Orientation: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.OriginResponseBody)
	r, _ := json.Marshal(res)
	t.Log(string(r))
}

func TestUsers_Likes(t *testing.T) {
	userGroup := Users{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := userGroup.Likes(UserLikesReq{
		Username:    "br00m",
		Page:        "1",
		PerPage:     "1",
		OrderBy:     "",
		Orientation: "",
	})
	t.Log(response.requestURL)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", res)
	t.Logf("%#v", response)
	t.Log(response.OriginResponseBody)
}

func TestUsers_Collections(t *testing.T) {
	userGroup := Users{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := userGroup.Collections(UserCollectionsReq{
		Username: "br00m",
		Page:     "1",
		PerPage:  "1",
	})
	t.Log(response.requestURL)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}

func TestUsers_Statistics(t *testing.T) {
	userGroup := Users{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := userGroup.Statistics(UserStatisticsReq{
		Username: "br00m",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))

}
