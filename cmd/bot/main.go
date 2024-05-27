package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

var (
	API = "7044707180:AAGMAqNArCMBo7Vr_7ykq89gtd9qOnWSHDg"
)

func echo(b *gotgbot.Bot, ctx *ext.Context) error {
	gameURL := os.Getenv("front-end-url")
	if gameURL == "" {
		_, _ = b.SendMessage(ctx.EffectiveChat.Id, "Error on the server. Please try again", nil)
		return errors.New("back end game url is empty")
	}
	opts := &gotgbot.SendMessageOpts{
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "Open mini app", WebApp: &gotgbot.WebAppInfo{Url: fmt.Sprintf("%s/?user_id=%d", gameURL, ctx.EffectiveUser.Id)}},
			}},
		},
	}
	_, err := b.SendMessage(ctx.EffectiveChat.Id, "Game", opts)
	return err
}

func main() {
	bot, err := gotgbot.NewBot(API, nil)
	if err != nil {
		log.Panic(err)
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)
	dispatcher.AddHandler(handlers.NewMessage(message.Text, echo))

	// Start receiving updates.
	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", bot.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()

	port := "8080"
	if port == "" {
		port = "3000"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
