package chatgpt

import (
	"encoding/json"
)

func ChatGPTAPIController(message string) string {
	// ChatGPT 서비스 호출
	nc := NewChatGPT()

	// ChatGPT API 실행
	resp := nc.ChatGPTAPI(message)

	// 응답 결과 JSON 문자열로 변환
	b, _ := json.Marshal(resp)

	// 결과 반환
	return string(b)
}
