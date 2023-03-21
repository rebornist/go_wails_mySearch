package gsearch

import "os"

func NewGSearch() *GSearchRequest {
	return &GSearchRequest{
		URL: "https://www.googleapis.com/customsearch/v1",
		Key: os.Getenv("GSearch_API_KEY"),
		CX:  os.Getenv("GSearch_API_CX"),
	}
}
