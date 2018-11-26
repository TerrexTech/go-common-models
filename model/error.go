package model

// Error describes any error.
type Error struct {
	Code    int16  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// NewError is a convenient awy to create an Error-instance.
func NewError(code int16, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}
