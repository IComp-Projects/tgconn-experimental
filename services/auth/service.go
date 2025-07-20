package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/icomp-projects/tgconn-experimental/shared"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type service struct {
	botToken        string
	sessionDuration time.Duration
}

func NewService(botToken string, sessionDuration time.Duration) *service {
	return &service{
		botToken:        botToken,
		sessionDuration: sessionDuration,
	}
}

func (s *service) ValidateInitData(w http.ResponseWriter, r *http.Request) {
	authParts := strings.Split(r.Header.Get("Authorization"), " ")

	if len(authParts) != 2 || authParts[0] != "tma" {
		http.Error(w, `{"message": "unauthorized"}`, http.StatusUnauthorized)
		return
	}

	authData := authParts[1]

	err := initdata.Validate(authData, s.botToken, s.sessionDuration)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	initData, err := initdata.Parse(authData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.WriteJSON(w, http.StatusOK, initData)
}
