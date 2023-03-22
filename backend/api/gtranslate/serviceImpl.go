package gtranslate

import (
	"changeme/backend/api/dto"
	res "changeme/backend/api/web/response"
	"encoding/json"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func (ctx *GTranslateRequest) GTranslateAPI(langText, text string) res.Response {

	var responseData []string

	// 번역 대상 국가 파싱
	lang, err := language.Parse(langText)
	if err != nil {
		return res.CommErrorResponse(err)
	}

	// 번역 클라이언트 생성
	client, err := translate.NewClient(ctx.Ctx)
	if err != nil {
		return res.CommErrorResponse(err)
	}
	defer client.Close()

	// 번역 실행
	resp, err := client.Translate(ctx.Ctx, []string{text}, lang, nil)
	if err != nil {
		return res.CommErrorResponse(err)
	}
	if len(resp) == 0 {
		return res.CommErrorResponse(err)
	}

	for _, item := range resp {

		var dto dto.SearchDto
		dto.Title = text
		dto.Content = item.Text

		// JSON 변환된 API 응답을 출력합니다.
		jsonStr, err := json.Marshal(dto)
		if err != nil {
			return res.CommErrorResponse(err)
		}

		responseData = append(responseData, string(jsonStr))
	}

	// 번역 결과 반환
	return res.CommSuccessResponse(responseData)
}
