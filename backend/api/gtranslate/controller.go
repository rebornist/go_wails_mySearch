package gtranslate

import (
	"encoding/json"
)

func GTranslateAPIController(targetLanguage, text string) string {
	// GTranslate 서비스 호출
	nc := NewGTranslate()

	// GTranslate API 실행
	resp := nc.GTranslateAPI(targetLanguage, text)

	// 응답 결과 JSON 문자열로 변환
	b, _ := json.Marshal(resp)

	// 결과 반환
	return string(b)
}
