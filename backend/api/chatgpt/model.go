package chatgpt

type ChatGPTRequest struct {
	URL string `json:"url"`
	Key string `json:"key"`
}

type ChatGPTResponse struct {
	Id      string          `json:"id"`
	Object  string          `json:"object"`
	Created int             `json:"created"`
	Choices []ChatGPTChoice `json:"choices"`
	Usage   ChatGPTUsage    `json:"usage"`
}

type ChatGPTChoice struct {
	Index        int                  `json:"index"`
	Message      ChatGPTChoiceMessage `json:"message"`
	FinishReason string               `json:"finish_reason"`
}

type ChatGPTChoiceMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
