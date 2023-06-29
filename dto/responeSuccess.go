package dto

type ResponeSuccess struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func WebResponeSuccess(code int, status string, data interface{}) ResponeSuccess {
	respone := ResponeSuccess{
		Code:   code,
		Status: status,
		Data:   data,
	}

	return respone
}
