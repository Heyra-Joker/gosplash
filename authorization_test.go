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

package main

import (
	"testing"
)

func TestClient_UserAuthentication(t *testing.T) {
	a := Authentication{Client: NewClient(accessKey)}
	authURL := a.UserAuthentication(UserAuthenticationReq{
		RedirectURI: "urn:ietf:wg:oauth:2.0:oob",
		Scope:       "public+read_user+write_user+write_likes",
	})

	t.Log(authURL)
}

func TestAuthentication_AuthenticationToken(t *testing.T) {
	a := Authentication{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	token, response, err := a.AuthenticationToken(AuthenticationTokenReq{
		ClientSecret: "8BIWgR3XtlTGGAtRtZZI3AY3j99mcu6LVBjDSNZK8RY",
		RedirectURI:  "urn:ietf:wg:oauth:2.0:oob",
		Code:         "WLDDPnH-3sHHiOIX0yxdFYv9TeBNfcP7bs85i3kRbl4",
		GrantType:    "authorization_code",
		ClientID:     accessKey,
	})
	t.Logf("%#v", response)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", token)
}
