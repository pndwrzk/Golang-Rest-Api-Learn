package dto

type Respone struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	ErrorMessage string `json:"error_message"`
}

func WebRespone(code int, status string, data interface{}, errorMessage string) Respone {
	respone := Respone{
		Code:   code,
		Status: status,
		Data:   data,
		ErrorMessage:errorMessage ,
	}

	return respone
}
