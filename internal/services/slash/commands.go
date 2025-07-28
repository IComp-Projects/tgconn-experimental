package slash

import (
	"context"
	"time"

	"github.com/icomp-projects/tgconn/internal/env"
	"github.com/icomp-projects/tgconn/internal/helpers"
	"github.com/icomp-projects/tgconn/internal/types"
	"github.com/imroc/req/v3"
)

var (
	BACKEND_BASE_URL = env.GetString(
		"BACKEND_BASE_URL",
		"https://bot-telegram-test-server1.onrender.com")

	client = req.C().SetTimeout(5 * time.Second)
	// instead of DevMode could be SetTimeout(5 * time.Second)
)

func Start(_ context.Context) types.Message {

	criarEnquete := *helpers.MakeWebAppButton("📝 Criar enquete", "/createPolls")
	vincularGrupo := *helpers.MakeWebAppButton("🔗 Vincular grupo", "/vincularGrupo")
	ajuda := *helpers.MakeWebAppButton("❓ Ajuda", "/help")

	return types.Message{
		Text: "Jax 🤖 na área. \n\nO que deseja fazer?",
		ReplyMarkup: &types.ReplyMarkup{
			InlineKeyboard: [][]types.InlineKeyboardButton{
				{criarEnquete},
				{vincularGrupo},
				{ajuda},
			},
		},
	}
}

func Bind(ctx context.Context, in types.BindInput) types.Message {
	m := types.Message{
		Text: "❌ Não foi possível vincular este grupo.\nTente novamente ou contate o suporte.",
	}

	resp, err := client.R().
		SetBody(in).
		Post(BACKEND_BASE_URL + "/api/bind-group/")

	if err != nil || !resp.IsSuccessState() {
		return m
	}

	m.Text = "✅ Grupo vinculado com sucesso!"
	return m
}
