package response

import "net/http"

func (res *Response) ErrorResponse(err error) Response {
	res.Code = http.StatusInternalServerError
	res.Message = err.Error()
	res.Data = nil
	return *res
}

func (res *Response) SuccessResponse(data []string) Response {
	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data
	return *res
}
