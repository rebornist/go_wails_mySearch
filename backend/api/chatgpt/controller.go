package chatgpt

import "fmt"

func ChatGPTAPIController(message string) string {
	// ChatGPT 서비스 호출
	nc := NewChatGPT()

	// ChatGPT API 실행
	respTxt, err := nc.ChatGPTAPI(message)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}

	// 결과 반환
	return fmt.Sprintf("%s", respTxt)
}
