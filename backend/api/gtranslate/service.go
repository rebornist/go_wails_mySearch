package gtranslate

type GTranslateService interface {
	GTranslateAPI(string, string) (string, error)
}
