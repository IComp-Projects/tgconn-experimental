package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/icomp-projects/tgconn-experimental/shared"
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

	if err := shared.ReadJSON(r, &update); err != nil {
		http.Error(w, fmt.Sprintf("Error reading JSON: %s", err), http.StatusBadRequest)
		return
	}

	log.Println(update)

	var msg Message
	switch update.Message.Text {
	case "/start":
		msg = s.cmd.Start(r.Context())
	case "/bind":
		in := BindInput{
			TelegramID: update.Message.From.ID,
			ChatID:     update.Message.Chat.ID,
			ChatTitle:  update.Message.Chat.Title,
		}
		msg = s.cmd.Bind(r.Context(), in)
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

	shared.WriteJSON(w, http.StatusOK, reply)
}

func (s *server) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/webhook", s.handleUpdates)

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
