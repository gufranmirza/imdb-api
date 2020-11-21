package errorinterface

// ErrorResponse renderer type for handling all sorts of errors.
type ErrorResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	Status string `json:"satus,omitempty"` // user-level status message
	Code   int64  `json:"code,omitempty"`  // application-specific error code
	Error  string `json:"error,omitempty"` // application-level error message, for debugging
}
