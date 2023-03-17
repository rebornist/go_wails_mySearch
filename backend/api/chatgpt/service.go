package chatgpt

type ChatGPTService interface {
	ChatGPTAPI(string) (string, error)
}
