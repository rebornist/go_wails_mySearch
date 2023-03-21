package gsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (g *GSearchRequest) GSearchAPI(message string) (string, error) {

	var responseData GSearchResponse

	// API 요청을 생성
	req, err := http.NewRequest("GET", g.URL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
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

	// API 응답 중 필요한 부분만 추출하여 JSON 변환.
	var respJson []map[string]string
	for _, item := range responseData.Items {
		respJsonItem := make(map[string]string)
		respJsonItem["title"] = item.Title
		respJsonItem["link"] = item.Link
		respJsonItem["snippet"] = item.Snippet
		respJson = append(respJson, respJsonItem)
	}

	// JSON 변환된 API 응답을 출력합니다.
	jsonStr, err := json.Marshal(respJson)
	if err != nil {
		fmt.Println("Error marshalling response:", err)
		return "", err
	}

	return string(jsonStr), nil
}
