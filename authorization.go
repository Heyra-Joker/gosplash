package main

import (
	"github.com/google/go-querystring/query"
	"net/url"
)

// Authentication Groups
type Authentication struct {
	Client *Client
}

// UserAuthenticationReq
// get user-authentication url request params
type UserAuthenticationReq struct {
	// how to find my redirect uri ?
	// find you application in `your apps` button then you can see part of `Redirect URI & Permissions` (just search it)
	// in you application page
	RedirectURI string `url:"redirect_uri"`
	// find all scope url is: https://unsplash.com/documentation/user-authentication-workflow#permission-scopes
	Scope string `url:"scope"`
}

func (u UserAuthenticationReq) API() string {
	return "https://unsplash.com/oauth/authorize"
}

func (u UserAuthenticationReq) Params() map[string]string {
	return map[string]string{
		"response_type": "code",
	}
}

// UserAuthentication create user-authentication url
// documents: https://unsplash.com/documentation#user-authentication
func (a *Authentication) UserAuthentication(req UserAuthenticationReq) (authURL string) {
	v, _ := query.Values(req)
	v.Add("client_id", a.Client.clientId)
	for k, p := range req.Params() {
		v.Add(k, p)
	}
	q, _ := url.QueryUnescape(v.Encode())
	authURL = req.API() + "?" + q
	return

}

// AuthenticationTokenReq create token params
type AuthenticationTokenReq struct {
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	Code         string `json:"code"`
	GrantType    string `json:"grant_type"` // Value "authorization_code".
	ClientID     string `json:"client_id"`
}

// AuthenticationTokenReply .
type AuthenticationTokenReply struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	CreatedAt    string `json:"created_at"`
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
}

func (AuthenticationTokenReq) API() string {
	return "https://unsplash.com/oauth/token"
}

func (AuthenticationTokenReq) Params() map[string]string {
	return map[string]string{}
}

// AuthenticationToken get authentication token
// documents: https://unsplash.com/documentation/user-authentication-workflow#authorization-workflow
func (a *Authentication) AuthenticationToken(req Req) (reply *AuthenticationTokenReply, response *OmitResponse, err error) {
	response, err = a.Client.request("POST", req, &reply, nil)
	return
}
