package helpers

import (
	"github.com/icomp-projects/tgconn/internal/env"
	"github.com/icomp-projects/tgconn/internal/types"
)

var FRONTEND_BASE_URL = env.GetString(
	"FRONTEND_BASE_URL",
	"https://poll-miniapp.vercel.app")

func MakeWebAppButton(label string, path string) *types.InlineKeyboardButton {
	return &types.InlineKeyboardButton{
		Text: label,
		WebApp: &types.WebAppInfo{
			URL: FRONTEND_BASE_URL + path,
		},
	}
}
