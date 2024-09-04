package handlers

import (
	"tapp_server/cmd/bot"

	tele "gopkg.in/telebot.v3"
)

func NewChatHandler(bot *bot.Tapp) *Handler {
	return &Handler{
		Bot: bot,
	}
}

func (h *Handler) StartHandler(c tele.Context, menu tele.ReplyMarkup) error {
	return c.Send("тестовый апп", &menu)
}

func (h *Handler) MessageHandler(c tele.Context) error {
	return c.Send(c.Text())
}
