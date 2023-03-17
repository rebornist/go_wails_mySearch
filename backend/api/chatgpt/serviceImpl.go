package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (gpt *ChatGPTRequest) ChatGPTAPI(message string) (string, error) {
	var responseData ChatGPTResponse

	// API 호출에 필요한 데이터를 생성
	data := []byte(fmt.Sprintf(`{
        "model": "gpt-3.5-turbo",
    	"messages": [{"role": "user", "content": "%s"}]
    }`, message))

	// API 요청을 생성
	req, err := http.NewRequest("POST", gpt.URL, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+gpt.Key)

	// API 호출 실행
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error calling API:", err)
		return "", err
	}
	defer response.Body.Close()

	// API 응답 불러오기
	var buf bytes.Buffer
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	// API 응답을 JSON으로 파싱
	err = json.Unmarshal(buf.Bytes(), &responseData)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return "", err
	}

	for _, choice := range responseData.Choices {
		if choice.FinishReason == "stop" {
			return choice.Message.Content, nil
		}
	}

	// API 응답을 출력합니다.
	return "", fmt.Errorf("no response")
}
