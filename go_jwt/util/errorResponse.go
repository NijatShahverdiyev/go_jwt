package util

type ErrorResponse struct {
	Status  bool
	Message string
}

func (e *ErrorResponse) ErrorMessage(status bool, message string) ErrorResponse {
	e.Status = status
	e.Message = message
	return ErrorResponse{Status: e.Status, Message: e.Message}
}
