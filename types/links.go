package types

type Links struct {
	Self             string `json:"self"`
	HTML             string `json:"html"`
	Photos           string `json:"photos"`
	Likes            string `json:"likes"`
	Portfolio        string `json:"portfolio"`
	Following        string `json:"following"`
	Followers        string `json:"followers"`
	Download         string `json:"download"`
	DownloadLocation string `json:"download_location"`
}
