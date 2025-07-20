package main

import (
  "log"
  "os"
)

func main () {
  srv := server{
    addr: ":5555",
    cmd: service{
      frontendURL: os.Getenv("FRONTEND_BASE_URL"),
      backendURL: os.Getenv("BACKEND_BASE_URL"),
    },
  }

  err := srv.Run()

  if err != nil {
    log.Fatalf("server failed to start: %v", err)
  }
}
