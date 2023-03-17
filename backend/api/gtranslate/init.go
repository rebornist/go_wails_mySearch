package gtranslate

import (
	"context"
)

func NewGTranslate() *GTranslateRequest {
	return &GTranslateRequest{
		Ctx: context.Background(),
	}
}
