package slash

import (
	"context"

	"github.com/icomp-projects/tgconn/internal/env"
	"github.com/icomp-projects/tgconn/internal/types"
	"github.com/imroc/req/v3"
)

var FRONTEND_BASE_URL = env.GetString(
	"FRONTEND_BASE_URL",
	"https://poll-miniapp.vercel.app")

var BACKEND_BASE_URL = env.GetString(
	"BACKEND_BASE_URL",
	"https://bot-telegram-test-server1.onrender.com")

func Start(_ context.Context) types.Message {
	criarEnquete := types.InlineKeyboardButton{
		Text: "Criar enquete",
		WebApp: &types.WebAppInfo{
			URL: FRONTEND_BASE_URL + "/createPolls",
		},
	}

	vincularGrupo := types.InlineKeyboardButton{
		Text: "Vincular grupo",
		WebApp: &types.WebAppInfo{
			URL: FRONTEND_BASE_URL + "/vinculo",
		},
	}

	return types.Message{
		Text: "Vamos começar 🖥️\nSelecione a opção desejada.",
		ReplyMarkup: &types.ReplyMarkup{
			InlineKeyboard: [][]types.InlineKeyboardButton{
				{criarEnquete},
				{vincularGrupo},
			},
		},
	}
}

func Bind(ctx context.Context, in types.BindInput) types.Message {
	m := types.Message{
		Text:        "Erro: Não foi possível vincular este grupo.",
		ReplyMarkup: nil,
	}

	client := req.C().DevMode()

	resp, err := client.R().
		SetBody(in).
		Post(BACKEND_BASE_URL + "/api/bind-group/")

	if err != nil || !resp.IsSuccessState() {
		return m
	}

	m.Text = "Sucesso: Grupo vinculado."
	return m
}
