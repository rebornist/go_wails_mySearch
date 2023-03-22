package response

func CommSuccessResponse(data []string) Response {
	res := NewResponse()
	return res.SuccessResponse(data)
}

func CommErrorResponse(err error) Response {
	res := NewResponse()
	return res.ErrorResponse(err)
}
