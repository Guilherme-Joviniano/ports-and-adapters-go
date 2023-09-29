package web_utils

type ErrorMessage struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func NewErrorJson(msg string, status int) *ErrorMessage {
	return &ErrorMessage{
		StatusCode: status,
		Message:    msg,
	}
}
