package internal

type Users struct {
	Client Client
}

type UserPublicProfileReq struct {
}

type UserPublicProfileReply struct {
}

func (p UserPublicProfileReq) API() string {
	return "https://api.unsplash.com/users/"
}

func (p UserPublicProfileReq) Params() map[string]string {
	return map[string]string{}
}

func (u *Users) PublicProfile(username string) (reply *UserPublicProfileReply, headers *ResponseHeaders, err error) {
	headers, err = u.Client.request("GET", UserPublicProfileReq{}, &reply, nil)
	return
}
