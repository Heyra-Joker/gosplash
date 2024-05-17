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

func TestTopics_TopicsList(t *testing.T) {
	topicsGroup := Topics{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := topicsGroup.TopicsList(TopicsListReq{
		IDs:     "",
		Page:    "1",
		PerPage: "1",
		OrderBy: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}

func TestTopics_GetATopic(t *testing.T) {
	topicsGroup := Topics{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := topicsGroup.GetATopic("ugc")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}

func TestTopics_TopicPhotos(t *testing.T) {
	topicsGroup := Topics{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := topicsGroup.TopicPhotos(TopicPhotosReq{
		IdOrSlug:    "B02qM02HN9g",
		Page:        "1",
		PerPage:     "1",
		Orientation: "",
		OrderBy:     "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}
