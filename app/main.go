package main

import (
	"log"
	"os"
	"strconv"
)

func main() {

	botID, err := strconv.ParseInt(os.Getenv("BOT_ID"), 10, 64)

	if err != nil {
		log.Fatalf("failed to parse bot id from environment: %v", err)
	}

	srv := server{
		addr: ":5555",
		cmd: service{
			frontendURL: os.Getenv("FRONTEND_BASE_URL"),
			backendURL:  os.Getenv("BACKEND_BASE_URL"),
		},
		botID: botID,
	}

	err = srv.Run()

	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
