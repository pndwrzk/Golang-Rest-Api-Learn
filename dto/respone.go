package dto

type Respone struct {
	Code         int         `json:"code"`
	Status       string      `json:"status"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error_message"`
}

type ResponeGetAll struct {
	Code         int         `json:"code"`
	Status       string      `json:"status"`
	Data         interface{} `json:"data"`
	Meta         interface{} `json:"meta"`
	ErrorMessage string      `json:"error_message"`
}

type ResultPaginate struct {
	Limit  int
	Offset int
	Order  string
}

func WebRespone(code int, status string, data interface{}, errorMessage string) Respone {
	respone := Respone{
		Code:         code,
		Status:       status,
		Data:         data,
		ErrorMessage: errorMessage,
	}

	return respone
}

func WebResponeGetAll(code int, status string, data interface{}, meta interface{}, errorMessage string) ResponeGetAll {
	respone := ResponeGetAll{
		Code:         code,
		Status:       status,
		Data:         data,
		Meta:         meta,
		ErrorMessage: errorMessage,
	}

	return respone
}
