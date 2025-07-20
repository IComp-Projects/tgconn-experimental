package main

import (
	"context"

	"github.com/imroc/req/v3"
)

type service struct {
	frontendURL string
	backendURL  string
}

func NewService(f string, b string) *service {
	return &service{
		frontendURL: f,
		backendURL:  b,
	}
}

func (s *service) Start(_ context.Context) Message {
	criarEnquete := InlineKeyboardButton{
		Text: "Criar enquete",
		WebApp: &WebAppInfo{
			URL: s.frontendURL + "/createPolls",
		},
	}

	vincularGrupo := InlineKeyboardButton{
		Text: "Vincular grupo",
		WebApp: &WebAppInfo{
			URL: s.frontendURL + "/vinculo",
		},
	}

	return Message{
		Text: "Vamos começar 🖥️\nSelecione a opção desejada.",
		ReplyMarkup: ReplyMarkup{
			InlineKeyboard: [][]InlineKeyboardButton{
				{criarEnquete},
				{vincularGrupo},
			},
		},
	}
}

func (s *service) Bind(ctx context.Context, in BindInput) Message {
	m := Message{
		Text:        "Erro: Não foi possível vincular este grupo.",
		ReplyMarkup: ReplyMarkup{},
	}

	client := req.C().DevMode()

	resp, err := client.R().
		SetBody(in).
		Post(s.backendURL + "/bind-group/")

	if err != nil || !resp.IsSuccessState() {
		return m
	}

	m.Text = "Sucesso: Grupo vinculado."
	return m
}
