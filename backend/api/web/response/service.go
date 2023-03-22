package response

type ResponseService interface {
	ErrorResponse(error) Response
	SuccessResponse(data []string) Response
}
