package gsearch

import (
	"bytes"
	"changeme/backend/api/dto"
	res "changeme/backend/api/web/response"
	"encoding/json"
	"net/http"
)

func (g *GSearchRequest) GSearchAPI(message string) res.Response {
	var responseData GSearchResponse

	// API 요청을 생성
	req, err := http.NewRequest("GET", g.URL, nil)
	if err != nil {
		return res.CommErrorResponse(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// API 호출에 필요한 파라메터를 생성
	q := req.URL.Query()
	q.Add("q", message)
	q.Add("num", "10")
	q.Add("start", "1")
	q.Add("key", g.Key)
	q.Add("cx", g.CX)

	req.URL.RawQuery = q.Encode()

	// API 호출 실행
	client := &http.Client{}
	respApi, err := client.Do(req)
	if err != nil {
		return res.CommErrorResponse(err)
	}
	defer respApi.Body.Close()

	// API 응답 불러오기
	var buf bytes.Buffer
	_, err = buf.ReadFrom(respApi.Body)
	if err != nil {
		return res.CommErrorResponse(err)
	}

	// API 응답을 JSON으로 파싱
	err = json.Unmarshal(buf.Bytes(), &responseData)
	if err != nil {
		return res.CommErrorResponse(err)
	}

	// API 응답 중 필요한 부분만 추출하여 []string 변환.
	var respJson []string
	for _, item := range responseData.Items {

		var dto dto.SearchDto
		dto.Title = item.Title
		dto.Link = item.Link
		dto.Content = item.Snippet

		// JSON 변환된 API 응답을 출력합니다.
		jsonStr, err := json.Marshal(dto)
		if err != nil {
			return res.CommErrorResponse(err)
		}

		respJson = append(respJson, string(jsonStr))
	}

	return res.CommSuccessResponse(respJson)
}
