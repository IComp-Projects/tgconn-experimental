package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type server struct {
	addr string
	cmd  service
}

func NewServer(addr string, svc service) *server {
	return &server{
		addr: addr,
		cmd:  svc,
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

func (s *server) Run() error {
	router := http.NewServeMux()

	authExpiry := time.Hour
	botID, err := strconv.ParseInt(os.Getenv("BOT_ID"), 10, 64)

	if err != nil {
		panic(err)
	}

	router.HandleFunc("/webhook", s.handleUpdates)
	router.Handle("/auth", CORS(s.cmd.frontendURL, Auth(botID, authExpiry, nil)))

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
