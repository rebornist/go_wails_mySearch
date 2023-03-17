package gtranslate

import "fmt"

func GTranslateAPIController(targetLanguage, text string) string {
	// GTranslate 서비스 호출
	nc := NewGTranslate()

	// GTranslate API 실행
	respTxt, err := nc.GTranslateAPI(targetLanguage, text)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}

	// 결과 반환
	return fmt.Sprintf("%s", respTxt)
}
