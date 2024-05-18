package config

// Generic struct for HTTP response with message only
type MessageResponse struct {
	Message string `json:"message"`
}

// Generic struct for HTTP response with message and data
type MessageAndDataResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
