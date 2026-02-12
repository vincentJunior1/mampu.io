package response

import (
	"core-system/core/entities/response"
)

func NewResponse(statusCode int, msg string, data interface{}) response.Response {
	res := response.Response{
		StatusCode: statusCode,
		Message:    msg,
		Data:       data,
	}
	return res
}

func NewResponseWithError(statusCode int, msg string) response.Response {
	res := response.Response{
		StatusCode: statusCode,
		Message:    msg,
	}
	return res
}
