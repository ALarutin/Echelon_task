package handlers

type Request struct {
	Sites      []string `json:"sites"`
	SearchText string   `json:"search-text"`
}

type Response struct {
	FoundAtSites []string `json:"found-at-sites"`
}
