package gsearch

type GSearchService interface {
	GSearchAPI(message string) (string, error)
}
