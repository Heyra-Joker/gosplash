package types

// AlternativeSlugs mismatch some value
// more detail see: https://unsplash.com/documentation#supported-languages
type AlternativeSlugs struct {
	En string `json:"en"`
	Es string `json:"es"`
	Ja string `json:"ja"`
	Fr string `json:"fr"`
	It string `json:"it"`
	Ko string `json:"ko"`
	De string `json:"de"`
	Pt string `json:"pt"`
}
