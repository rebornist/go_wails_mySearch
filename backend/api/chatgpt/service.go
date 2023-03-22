package chatgpt

import (
	"changeme/backend/api/web/response"
)

type ChatGPTService interface {
	ChatGPTAPI(string) response.Response
}
