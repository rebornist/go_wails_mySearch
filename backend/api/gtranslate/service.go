package gtranslate

import (
	"changeme/backend/api/web/response"
)

type GTranslateService interface {
	GTranslateAPI(string, string) (response.Response, error)
}
