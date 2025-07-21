package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func CORS(allowedOrigin string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("origin")

		if origin == allowedOrigin {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Auth(botID int64, expiry time.Duration, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authParts := strings.Split(r.Header.Get("Authorization"), " ")

		if len(authParts) != 2 || authParts[0] != "tma" {
			http.Error(w, `{"message": "unauthorized"}`, http.StatusUnauthorized)
			return
		}

		authData := authParts[1]

		err := initdata.ValidateThirdParty(authData, botID, expiry)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			log.Println(authData, err.Error())
			return
		}

		initData, err := initdata.Parse(authData)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		WriteJSON(w, http.StatusOK, initData)

		next.ServeHTTP(w, r)
	})
}
