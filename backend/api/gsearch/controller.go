package gsearch

import (
	"encoding/json"
)

func GSearchController(message string) string {
	// GSearch 서비스 호출
	gs := NewGSearch()

	// GSearch API 실행
	resp := gs.GSearchAPI(message)

	// 응답 결과 JSON 문자열로 변환
	b, _ := json.Marshal(resp)

	// 결과 반환
	return string(b)
}
