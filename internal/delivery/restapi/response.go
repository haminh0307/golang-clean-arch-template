package restapi

type Response struct {
	// Data contains the response body for the success case.
	Data any `json:"data,omitempty"`
	// In other cases, this contains the error.
	Error string `json:"error,omitempty"`
}
