package gsearch

import (
	"changeme/backend/api/web/response"
)

type GSearchService interface {
	GSearchAPI(message string) response.Response
}
