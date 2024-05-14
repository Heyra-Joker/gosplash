package types

type Historical struct {
	Change     int64  `json:"change"`
	Average    int64  `json:"average"`
	Resolution string `json:"resolution"`
	Quantity   int64  `json:"quantity"`
}
