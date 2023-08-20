package helpers

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Subdistrict struct {
	Data interface{} `json:"subdistrict"`
}

type Village struct {
	Data interface{} `json:"village"`
}

func APIResponse(message string, code int, data interface{}) Response {

	jsonResponse := Response{
		Status:  code,
		Message: message,
		Data:    data,
	}

	return jsonResponse
}

func APIResponseSubdistrict(message string, code int, data interface{}) Response {

	jsonResponse := Response{
		Status:  code,
		Message: message,
		Data: Subdistrict{
			Data: data,
		},
	}

	return jsonResponse
}

func APIResponseVillage(message string, code int, data interface{}) Response {

	jsonResponse := Response{
		Status:  code,
		Message: message,
		Data: Village{
			Data: data,
		},
	}

	return jsonResponse
}
