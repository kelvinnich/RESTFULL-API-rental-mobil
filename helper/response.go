package helper

type Response struct{
	Status bool
	Message string
	Errors interface{}
	Data interface{}
}

type EmptyObj struct {}

func ResponseOK(status bool, message string, data interface{}) Response{
	res := Response{
		Status: status,
		Message: message,
		Errors: nil,
		Data: data,
	}

	return res
}

func ResponseERROR( message string, err interface{}, data interface{}) Response{
	res := Response{
		Status: false,
		Message: message,
		Errors: err,
		Data: data,
	}
	return res
}