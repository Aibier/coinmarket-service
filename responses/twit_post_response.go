package responses

// TwitPostResponse ...
type TwitPostResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// RecentSearchAPIResponse ...
type RecentSearchAPIResponse struct {
	Data []interface{} `json:"data"`
	Meta Meta `json:"meta"`
}

// Meta ...
type Meta struct {
	NewestID string `json:"newest_id"`
	OldestID string `json:"oldest_id"`
	ResultCount int `json:"result_count"`
	NextToken string `json:"next_token"`
}