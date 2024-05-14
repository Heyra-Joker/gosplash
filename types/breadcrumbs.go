package types

type Breadcrumbs struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Index int    `json:"index"`
	Type  string `json:"type"`
}
