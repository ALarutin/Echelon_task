package handlers

type request struct {
	Sites      []string `json:"sites"`
	SearchText string   `json:"search-text"`
}

type response struct {
	FoundAtSites []string `json:"found-at-sites"`
}

type errorJSON struct {
	Error string `json:"error"`
}
