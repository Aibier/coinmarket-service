package responses

// ErrorResponse ...
type ErrorResponse struct {
	Status  int                    `json:"status"`
	Detail string                 `json:"detail"`
	Type string                 `json:"type"`
	Title    map[string]interface{} `json:"title"`
}