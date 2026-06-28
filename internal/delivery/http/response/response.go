package response

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Success(data any, message string) Response {
	return Response{
		Success: true,
		Data:    data,
		Message: message,
	}
}

func Failure(err string, message string) Response {
	return Response{
		Success: false,
		Error:   err,
		Message: message,
	}
}
