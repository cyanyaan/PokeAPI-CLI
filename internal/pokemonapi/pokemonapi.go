package pokemonapi

type Client struct {
	http
}

type LocationAreasResp struct {
	Count    int     `json:"count,omitempty"`
	Next     *string `json:"next,omitempty"`
	Previous *any    `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"results,omitempty"`
}
