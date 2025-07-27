package api

import (
	"fmt"
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

	var msg types.Message
	switch update.Message.Text {
	case "/start":
		msg = cli.Start(r.Context())
	case "/bind":
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
