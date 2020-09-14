package response

type response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func NewResponse(code int, message string, data interface{}) response {
	return response{Code: code, Message: message, Data: data}
}
