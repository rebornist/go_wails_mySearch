package chatgpt

import (
	"bytes"
	"changeme/backend/api/dto"
	res "changeme/backend/api/web/response"
	"encoding/json"
	"fmt"
	"net/http"
)

func (gpt *ChatGPTRequest) ChatGPTAPI(message string) res.Response {
	var getResponse ChatGPTResponse
	var responeData []string

	// API 호출에 필요한 데이터를 생성
	data := []byte(fmt.Sprintf(`{
        "model": "gpt-3.5-turbo",
    	"messages": [{"role": "user", "content": "%s"}]
    }`, message))

	// API 요청을 생성
	req, err := http.NewRequest("POST", gpt.URL, bytes.NewBuffer(data))
	if err != nil {
		return res.CommErrorResponse(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+gpt.Key)

	// API 호출 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res.CommErrorResponse(err)
	}
	defer resp.Body.Close()

	// API 응답 불러오기
	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return res.CommErrorResponse(err)
	}

	// API 응답을 JSON으로 파싱
	err = json.Unmarshal(buf.Bytes(), &getResponse)
	if err != nil {
		return res.CommErrorResponse(err)
	}

	for _, choice := range getResponse.Choices {
		if choice.FinishReason == "stop" {
			var dto dto.SearchDto
			dto.Title = message
			dto.Content = choice.Message.Content

			// JSON 변환된 API 응답을 출력합니다.
			jsonStr, err := json.Marshal(dto)
			if err != nil {
				return res.CommErrorResponse(err)
			}

			// API 응답 중 필요한 부분만 추출하여 JSON 변환.
			responeData = append(responeData, string(jsonStr))
		}
	}

	// API 응답을 출력합니다.
	return res.CommSuccessResponse(responeData)
}
