package chatgpt

import "os"

func NewChatGPT() *ChatGPTRequest {
	return &ChatGPTRequest{
		URL: "https://api.openai.com/v1/chat/completions",
		Key: os.Getenv("ChatGPT_API_KEY"),
	}
}
