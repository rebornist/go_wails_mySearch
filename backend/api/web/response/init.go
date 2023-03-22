package response

func NewResponse() Response {
	return Response{
		Code:    -1,
		Message: "",
		Data:    nil,
	}
}
