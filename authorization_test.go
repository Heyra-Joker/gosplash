package main

import (
	"testing"
)

func TestClient_UserAuthentication(t *testing.T) {
	a := Authentication{Client: NewClient(accessKey)}
	authURL := a.UserAuthentication(UserAuthenticationReq{
		RedirectURI: "urn:ietf:wg:oauth:2.0:oob",
		Scope:       "public+read_user+write_user",
	})

	t.Log(authURL)
}

func TestAuthentication_AuthenticationToken(t *testing.T) {
	a := Authentication{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	token, response, err := a.AuthenticationToken(AuthenticationTokenReq{
		ClientSecret: "8BIWgR3XtlTGGAtRtZZI3AY3j99mcu6LVBjDSNZK8RY",
		RedirectURI:  "urn:ietf:wg:oauth:2.0:oob",
		Code:         "k3RBzJzg-xpppflGJ7uQMtdDsLpmawhlZbnsSWPNCsQ",
		GrantType:    "authorization_code",
		ClientID:     accessKey,
	})
	t.Logf("%#v", response)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", token)
}
