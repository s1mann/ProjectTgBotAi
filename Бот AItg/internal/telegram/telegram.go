package telegram

import (
	ai "app/AI"
	f "fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run(token string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = false

	f.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			f.Printf("user: %s, msg: %s\n", update.Message.From.UserName, update.Message.Text)

			if update.Message.Text == "/start" {

				go ai.Text(bot,
					update.Message.Chat.ID,
					f.Sprintf("Привет! Меня зовут %s!", update.Message.From.UserName),
				)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Хорошо, секунду...")
				bot.Send(msg)
			} else {

				go ai.Text(bot, update.Message.Chat.ID, update.Message.Text)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Хорошо, секунду...")
				bot.Send(msg)
			}
		}
	}
	return nil
}
