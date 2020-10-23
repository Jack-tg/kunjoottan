package main

import (
	"fmt"
	"log"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbot.NewBotAPI("1393766536:AAEoRyXgsNVfrAQMy6k1Y1xtmP6wAyaj694") //token from botfather
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Print("connected duccessfully")

	u := tgbot.NewUpdate(0)
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		
	
		msg := tgbot.NewMessage(update.Message.Chat.ID, "പോടീ പൂറി🤬🤬🤬🤬") //abbusive message
		bot.Send(msg)
	}
}
