package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/icomp-projects/tgconn/internal/helpers"
	"github.com/icomp-projects/tgconn/internal/services/cli"
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
		msg = cli.Start(r.Context())
	case "/bind":
		if update.Message.Chat.Type != "group" {
			msg = types.Message{
				Text:        "Erro: Comando /bind só pode ser usado em grupos.",
				ReplyMarkup: types.ReplyMarkup{},
			}
			break
		}
		in := update.AsBindInput()
		msg = cli.Bind(r.Context(), *in)
	default:
		msg = types.Message{
			Text:        "Erro: Comando não suportado.",
			ReplyMarkup: types.ReplyMarkup{},
		}
	}

	reply := types.TelegramUpdateReply{
		Method:      "sendMessage",
		ChatID:      update.Message.Chat.ID,
		Text:        msg.Text,
		ReplyMarkup: msg.ReplyMarkup,
	}

	helpers.WriteJSON(w, http.StatusOK, reply)
}
