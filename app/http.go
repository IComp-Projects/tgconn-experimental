package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type server struct {
	addr  string
	cmd   service
	botID int64
}

func NewServer(addr string, svc service, botID int64) *server {
	return &server{
		addr:  addr,
		cmd:   svc,
		botID: botID,
	}
}

func (s *server) handleUpdates(w http.ResponseWriter, r *http.Request) {
	var update TelegramUpdate

	if err := ReadJSON(r, &update); err != nil {
		http.Error(w, fmt.Sprintf("Error reading JSON: %s", err), http.StatusBadRequest)
		return
	}

	log.Println(update)

	var msg Message
	switch update.Message.Text {
	case "/start":
		msg = s.cmd.Start(r.Context())
	case "/bind":
		in := update.asBindInput()
		msg = s.cmd.Bind(r.Context(), *in)
	default:
		msg = Message{
			Text:        "Erro: Comando não suportado.",
			ReplyMarkup: ReplyMarkup{},
		}
	}

	reply := TelegramUpdateReply{
		Method:      "sendMessage",
		ChatID:      update.Message.Chat.ID,
		Text:        msg.Text,
		ReplyMarkup: msg.ReplyMarkup,
	}

	WriteJSON(w, http.StatusOK, reply)
}

func (s *server) handleAuth(w http.ResponseWriter, r *http.Request) {
	authParts := strings.Split(r.Header.Get("Authorization"), " ")

	if len(authParts) != 2 || authParts[0] != "tma" {
		http.Error(w, `{"message": "unauthorized"}`, http.StatusUnauthorized)
		return
	}

	authData := authParts[1]

	err := initdata.ValidateThirdParty(authData, s.botID, time.Hour)

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
}

func (s *server) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/webhook", s.handleUpdates)
	router.Handle("/auth", CORS(s.cmd.frontendURL, http.HandlerFunc(s.handleAuth)))

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
