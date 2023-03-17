package main

import (
	"context"
	"fmt"

	"changeme/backend/api/chatgpt"
	"changeme/backend/api/gtranslate"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// ChatGPT API
func (a *App) ChatGPT(message string) string {
	return chatgpt.ChatGPTAPIController(message)
}

// GoogleTranslate Google 번역기 API
func (a *App) GoogleTranslate(targetLanguage, text string) string {
	return gtranslate.GTranslateAPIController(targetLanguage, text)
}
