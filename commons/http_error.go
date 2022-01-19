package commons

type HTTPError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *HTTPError) Error() string {
	return e.Code + ": " + e.Message
}

func NewHTTPError(code string, msg string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: msg,
	}
}
