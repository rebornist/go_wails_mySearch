package gsearch

type GSearchRequest struct {
	URL string `json:"url"`
	Key string `json:"key"`
	CX  string `json:"cx"`
}

type GSearchResponse struct {
	Kind              string                           `json:"kind"`
	URL               GSearchResponseURL               `json:"url"`
	Queries           GSearchResponseQueries           `json:"queries"`
	Context           GSearchResponseContext           `json:"context"`
	SearchInformation GSearchResponseSearchInformation `json:"searchInformation"`
	Items             []GSearchResponseItem            `json:"items"`
	Spelling          GSearchResponseSpelling          `json:"spelling"`
}

type GSearchResponseURL struct {
	Type     string `json:"type"`
	Template string `json:"template"`
}

type GSearchResponseQueries struct {
	NextPage []GSearchResponseQueriesNextPage `json:"nextPage"`
	Request  []GSearchResponseQueriesRequest  `json:"request"`
}

type GSearchResponseQueriesNextPage struct {
	Title          string `json:"title"`
	TotalResults   string `json:"totalResults"`
	SearchTerms    string `json:"searchTerms"`
	Count          int    `json:"count"`
	StartIndex     int    `json:"startIndex"`
	InputEncoding  string `json:"inputEncoding"`
	OutputEncoding string `json:"outputEncoding"`
	Safe           string `json:"safe"`
	Cx             string `json:"cx"`
}

type GSearchResponseQueriesRequest struct {
	Title          string `json:"title"`
	TotalResults   string `json:"totalResults"`
	SearchTerms    string `json:"searchTerms"`
	Count          int    `json:"count"`
	StartIndex     int    `json:"startIndex"`
	InputEncoding  string `json:"inputEncoding"`
	OutputEncoding string `json:"outputEncoding"`
	Safe           string `json:"safe"`
	Cx             string `json:"cx"`
}

type GSearchResponseContext struct {
	Title string `json:"title"`
}

type GSearchResponseSearchInformation struct {
	SearchTime            float64 `json:"searchTime"`
	FormattedSearchTime   string  `json:"formattedSearchTime"`
	TotalResults          string  `json:"totalResults"`
	FormattedTotalResults string  `json:"formattedTotalResults"`
}

type GSearchResponseItems struct {
	Item []GSearchResponseItem
}

type GSearchResponseItem struct {
	Kind             string `json:"kind"`
	Title            string `json:"title"`
	HTMLTitle        string `json:"htmlTitle"`
	Link             string `json:"link"`
	DisplayLink      string `json:"displayLink"`
	Snippet          string `json:"snippet"`
	HTMLSnippet      string `json:"htmlSnippet"`
	CacheID          string `json:"cacheId"`
	FormattedURL     string `json:"formattedUrl"`
	HTMLFormattedURL string `json:"htmlFormattedUrl"`
}

type GSearchResponseItemsPagemap struct {
	CseThumbnail []GSearchResponseItemsPagemapCseThumbnail `json:"cse_thumbnail"`
	CseImage     []GSearchResponseItemsPagemapCseImage     `json:"cse_image"`
	Metatages    []GSearchResponseItemsPagemapMetatages    `json:"metatags"`
}

type GSearchResponseItemsPagemapCseThumbnail struct {
	Src    string `json:"src"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

type GSearchResponseItemsPagemapCseImage struct {
	Src string `json:"src"`
}

type GSearchResponseItemsPagemapMetatages struct {
	OgImage                         string `json:"og:image"`
	OgImageWidth                    string `json:"og:image:width"`
	OgImageHeight                   string `json:"og:image:height"`
	ThemeColor                      string `json:"theme-color"`
	OgType                          string `json:"og:type"`
	OgTitle                         string `json:"og:title"`
	OgDescription                   string `json:"og:description"`
	OgSiteName                      string `json:"og:site_name"`
	OgUrl                           string `json:"og:url"`
	TwitterCard                     string `json:"twitter:card"`
	TwitterTitle                    string `json:"twitter:title"`
	TwitterDescription              string `json:"twitter:description"`
	TwitterImage                    string `json:"twitter:image"`
	TwitterSite                     string `json:"twitter:site"`
	TwitterCreator                  string `json:"twitter:creator"`
	TiwtterDomain                   string `json:"tiwtter:domain"`
	Title                           string `json:"title"`
	Author                          string `json:"author"`
	Handheldfriendly                string `json:"handheldfriendly"`
	Google                          string `json:"google"`
	AppleMobileWebAppStatusBarStyle string `json:"apple-mobile-web-app-status-bar-style"`
	Viewport                        string `json:"viewport"`
	AppleMobileWebAppCapable        string `json:"apple-mobile-web-app-capable"`
	FormatDetection                 string `json:"format-detection"`
	Mobileoptimized                 string `json:"mobileoptimized"`
	MoblieWebAppCapable             string `json:"moblie-web-app-capable"`
}

type GSearchResponseSpelling struct {
	CorrectedQuery     string `json:"correctedQuery"`
	HTMLCorrectedQuery string `json:"htmlCorrectedQuery"`
}
