package stringsvc

import "errors"

// StringService provides operations on strings.
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("empty string")

// UppercaseRequest holds the request parameters for the Uppercase method.
type UppercaseRequest struct {
	S string `json:"s"`
}

// UppercaseResponse holds the response values for the Uppercase method.
type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

// CountRequest holds the request parameters for the Count method.
type CountRequest struct {
	S string `json:"s"`
}

// CountResponse holds the response values for the Count method.
type CountResponse struct {
	V int `json:"v"`
}
