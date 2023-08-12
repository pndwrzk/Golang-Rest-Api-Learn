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
	TotalData    int64       `json:"total_data"`
	Meta         interface{} `json:"meta"`
	ErrorMessage string      `json:"error_message"`
}

type ResponeLogin struct {
	Code         int    `json:"code"`
	Status       string `json:"status"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ResultPaginate struct {
	Limit  int
	Offset int
	Order  string
	Search string
}

func WebRespone(code int, status string, data interface{}, errorMessage string) Respone {

	var res Respone
	res.Code = code
	res.Status = status
	res.Data = data
	res.ErrorMessage = errorMessage

	return res
}

func WebResponeGetAll(code int, status string, data interface{}, totalData int64, meta interface{}, errorMessage string) ResponeGetAll {
	respone := ResponeGetAll{
		Code:         code,
		Status:       status,
		Data:         data,
		TotalData:    totalData,
		Meta:         meta,
		ErrorMessage: errorMessage,
	}

	return respone
}
