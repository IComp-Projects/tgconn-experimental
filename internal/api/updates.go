package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/icomp-projects/tgconn/internal/helpers"
	"github.com/icomp-projects/tgconn/internal/services/slash"
	"github.com/icomp-projects/tgconn/internal/types"
)

func (app *application) HandleUpdates(w http.ResponseWriter, r *http.Request) {
	var update types.TelegramUpdate

	if err := helpers.ReadJSON(r, &update); err != nil {
		http.Error(w, fmt.Sprintf("Error reading JSON: %s", err), http.StatusBadRequest)
		return
	}

	app.Logger.Info("incoming update",
		slog.Int64("from_id", update.Message.From.ID),
		slog.Int64("chat_id", update.Message.Chat.ID),
		slog.String("chat_title", update.Message.Chat.Title),
		slog.String("chat_type", update.Message.Chat.Type),
		slog.String("text", update.Message.Text),
	)

	var msg types.Message
	switch update.Message.Text {
	case "/start":
		msg = slash.Start(r.Context())
	case "/bind":
		if update.Message.Chat.Type != "group" {
			msg = types.Message{
				Text:        "O comando /bind só pode ser utilizado em grupos.",
				ReplyMarkup: nil,
			}
			break
		}
		in := update.AsBindInput()
		msg = slash.Bind(r.Context(), *in)
	default:
		msg = types.Message{
			Text:        "Erro: Comando não suportado.",
			ReplyMarkup: nil,
		}
	}

	reply := types.TelegramUpdateReply{
		Method: "sendMessage",
		ChatID: update.Message.Chat.ID,
		Text:   msg.Text,
	}

	if msg.ReplyMarkup != nil {
		reply.ReplyMarkup = msg.ReplyMarkup
	}

	helpers.WriteJSON(w, http.StatusOK, reply)
}
