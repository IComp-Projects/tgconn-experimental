package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	sessionDuration := time.Hour

	svc := NewService(botToken, sessionDuration)

	http.HandleFunc("/auth", svc.ValidateInitData)

	addr := ":4444"

	log.Println("Starting server on", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		panic(err)
	}
}
