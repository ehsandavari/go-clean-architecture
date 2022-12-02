package Common

type (
	ErrorResponse struct {
		Error ErrorDetails `json:"error"`
	}
	ErrorDetails struct {
		Message string `json:"message"`
	}
)

func errorResponse(err error) ErrorResponse {
	var message string
	if err != nil {
		message = err.Error()
	}
	return ErrorResponse{Error: ErrorDetails{
		Message: message,
	}}
}
