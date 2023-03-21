package gsearch

import "fmt"

func GSearchController(message string) string {
	// GSearch 서비스 호출
	gs := NewGSearch()

	// GSearch API 실행
	respTxt, err := gs.GSearchAPI(message)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}

	// 결과 반환
	return fmt.Sprintf("%s", respTxt)
}
