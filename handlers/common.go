package handlers

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func newResponse(status int, message string, result interface{}) *response {
	return &response{status, message, result}
}