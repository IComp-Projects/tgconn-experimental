package main

import (
	"net/http"
	"os"
	"time"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	sessionDuration := time.Hour

	svc := NewService(botToken, sessionDuration)

	http.HandleFunc("/auth", svc.ValidateInitData)

	err := http.ListenAndServe(":4444", nil)

	if err != nil {
		panic(err)
	}
}
