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
	"testing"
)

const accessKey = "44POH-LuAn-Knh-JrASbP2yGcDDKGDArtAAJhk9x_Lk"

func TestCurrentUser_Me(t *testing.T) {
	currentUserGroup := CurrentUser{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	reply, response, err := currentUserGroup.Me("ntMHq_GwOOdmkwVrk9oxch7nnaKZ29XrNSne56-dPjY")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", reply)
	t.Log(response.OriginResponseBody)
}

func TestCurrentUser_UpdateMe(t *testing.T) {
	currentUserGroup := CurrentUser{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	req := UpdateMeRequest{InstagramUsername: "Joker", FirstName: "thanks", Email: "77581236@gmail.com"}
	reply, response, err := currentUserGroup.UpdateMe("Kb7QKq_NAHxX5YiV713fKzWd3nFDMzKpQIs7hAYx47w", req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", reply)
	t.Logf(response.OriginResponseBody)
}
