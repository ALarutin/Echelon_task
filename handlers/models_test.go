package handlers_test

type successfulCase struct {
	number     uint64   `json:"-"`
	Sites      []string `json:"sites"`
	SearchText string   `json:"search-text"`
}
