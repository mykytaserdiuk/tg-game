package handler

import (
	"errors"
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func New(url string) *Handler {
	return &Handler{
		gameURL: url,
	}
}

type Handler struct {
	gameURL string
}

func (h *Handler) SendGame(b *gotgbot.Bot, ctx *ext.Context) error {
	if h.gameURL == "" {
		_, _ = b.SendMessage(ctx.EffectiveChat.Id, "Error on the server. Please try again", nil)
		return errors.New("front end game url is empty")
	}
	opts := &gotgbot.SendMessageOpts{
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "Open mini app", WebApp: &gotgbot.WebAppInfo{Url: fmt.Sprintf("%s/?user_id=%d", h.gameURL, ctx.EffectiveUser.Id)}},
			}},
		},
	}
	_, err := b.SendMessage(ctx.EffectiveChat.Id, "Game", opts)
	return err
}
