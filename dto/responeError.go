package dto

type ResponeError struct {
	Code         int    `json:"code"`
	Status       string `json:"status"`
	MessageError string `json:"message_error"`
}

func WebResponeError(code int, status string, message string) ResponeError {
	respone := ResponeError{
		Code:         code,
		Status:       status,
		MessageError: message,
	}

	return respone
}
