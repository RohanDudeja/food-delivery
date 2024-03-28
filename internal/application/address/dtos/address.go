package dtos

type (
	Address struct {
		Line1   string `json:"line1"`
		Line2   string `json:"line2"`
		City    string `json:"city"`
		State   string `json:"state"`
		Country string `json:"country"`
	}
)
