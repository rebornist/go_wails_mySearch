package gtranslate

import (
	"fmt"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func (ctx *GTranslateRequest) GTranslateAPI(langText, text string) (string, error) {

	// 번역 대상 국가 파싱
	lang, err := language.Parse(langText)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	// 번역 클라이언트 생성
	client, err := translate.NewClient(ctx.Ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	// 번역 실행
	resp, err := client.Translate(ctx.Ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}

	// 번역 결과 반환
	return resp[0].Text, nil
}
