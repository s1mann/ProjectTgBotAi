package ai

import (
	"context"
	f "fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	openai "github.com/sashabaranov/go-openai"
)

const token = "8109905612:AAHgYQW085wSPqJwhxSzWOHpoScdB8rEU4c"

func Text(bot *tgbotapi.BotAPI, id int64, msg string) {

	client := openai.NewClient(token)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		},
	)

	if err != nil {
		f.Printf("ChatComplete error: %v\n", err)
		return
	}

	message := tgbotapi.NewMessage(id, resp.Choices[0].Message.Content)
	bot.Send(message)

}
