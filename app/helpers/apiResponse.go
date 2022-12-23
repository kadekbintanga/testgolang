package helpers


type Response struct{
	Status	Status	`json:"status"`
	Data	any		`json:"data"`
}

type Status struct{
	Message	string	`json:"message"`
	Code	int		`json:"code"`
}

func APIResponse[D any](message string, code int, data D) Response{
	status := Status{
		Message: message,
		Code: code,
	}

	jsonResponse := Response{
		Status : status,
		Data : data,
	}

	return jsonResponse
}