package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/mykytaserdiuk/souptgbot/internal/bot/handler"
)

var (
	API = os.Getenv("TG_TOKEN")
)

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
	var gameURL = os.Getenv("front-end-url")
	handler := handler.New(gameURL)

	log.Println(os.Environ())
	updater := ext.NewUpdater(dispatcher, nil)
	dispatcher.AddHandler(handlers.NewMessage(message.Text, handler.SendGame))

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
