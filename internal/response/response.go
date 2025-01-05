package response

// Response http-server's response
type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	statusOK    = "OK"
	statusError = "Error"
)

// OK returns success response
func OK() Response {
	return Response{
		Status: statusOK,
	}
}

// Error returns error response
func Error(msg string) Response {
	return Response{
		Status: statusError,
		Error:  msg,
	}
}
