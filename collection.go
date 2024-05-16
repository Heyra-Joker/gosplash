package gosplash

type Collection struct {
	Client *Client
}

type CollectionReq struct {
	Page    string `json:"page,omitempty"`
	PerPage string `json:"per_page,omitempty"`
}

type CollectionReply struct {
}

func (CollectionReq) API() string {
	return "https://api.unsplash.com/collections"
}

func (CollectionReq) Params() map[string]string {
	return nil
}

func (c *Collection) Collections(req CollectionReq) (reply []*CollectionReply, response *OmitResponse, err error) {
	response, err = c.Client.request("GET", req, &reply, nil)
	return
}
