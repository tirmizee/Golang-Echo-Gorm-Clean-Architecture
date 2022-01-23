package commons

type CustomError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *CustomError) Error() string {
	return e.Code + ": " + e.Message
}

func NewCustomError(code string, msg string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: msg,
	}
}
