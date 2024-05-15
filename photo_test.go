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

func TestPhoto_Photos(t *testing.T) {
	photoGroup := Photo{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := photoGroup.Photos(PhotosReq{
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

func TestPhoto_GetPhoto(t *testing.T) {
	photoGroup := Photo{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := photoGroup.GetAPhoto("Nx2k4q0dcJY")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	//t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}

func TestPhoto_Random(t *testing.T) {
	photoGroup := Photo{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := photoGroup.Random(PhotoRandomReq{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}

func TestPhoto_Statistics(t *testing.T) {
	photoGroup := Photo{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := photoGroup.Statistics("a-tray-of-roasted-carrots-and-potatoes-KxUIupj29iU")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}

func TestPhoto_TrackDownload(t *testing.T) {
	photoGroup := Photo{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := photoGroup.TrackDownload("the-sun-is-setting-over-a-body-of-water-AhMoWv0fQzc")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}

func TestPhoto_Update(t *testing.T) {
	//photoGroup := Photo{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	//token := "ntMHq_GwOOdmkwVrk9oxch7nnaKZ29XrNSne56-dPjY"
	//res, response, err := photoGroup.Update(token, PhotoUpdateReq{
	//	ID:                  "",
	//	Description:         "",
	//	ShowOnProfile:       "",
	//	Tags:                "",
	//	LocationLatitude:    "",
	//	LocationLongitude:   "",
	//	LocationName:        "",
	//	LocationCity:        "",
	//	LocationCountry:     "",
	//	ExifMake:            "",
	//	ExifModel:           "",
	//	ExifExposureTime:    "",
	//	ExifApertureValue:   "",
	//	ExifFocalLength:     "",
	//	ExifIsoSpeedRatings: "",
	//})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(response.requestURL)
	//t.Log(response.OriginResponseBody)
	//a, _ := json.Marshal(res)
	//t.Logf(string(a))
}

func TestPhoto_Like(t *testing.T) {
	photoGroup := Photo{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	token := "PEfcJxijkWh2Nh2tCD-c24Mo-D9PiUVHrTCn1UjwfYI"
	res, response, err := photoGroup.Like(token, "a-tray-of-roasted-carrots-and-potatoes-KxUIupj29iU")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}

func TestPhoto_UnLike(t *testing.T) {
	photoGroup := Photo{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	token := "PEfcJxijkWh2Nh2tCD-c24Mo-D9PiUVHrTCn1UjwfYI"
	res, response, err := photoGroup.UnLike(token, "a-tray-of-roasted-carrots-and-potatoes-KxUIupj29iU")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}
