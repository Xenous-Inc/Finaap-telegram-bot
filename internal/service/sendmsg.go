package service

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var buttons = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Как скачать приложение"),
		tgbotapi.NewKeyboardButton("У меня проблема"),
	),
)

func SendMsg(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.Text == "/start" {
			msg.Text = "Привет, я бот финашки"
			msg.ReplyMarkup = buttons
		}

		if update.Message.Text == "Как скачать приложение" {
			msg.Text = "никак"
		}

		if update.Message.Text == "У меня проблема" {
			msg.Text = "Опишите проблему в одном сообщении начиная со слова 'Проблема:' Пример жалобы: Проблема: Не работает приложение"
		}

		words := strings.Split(update.Message.Text, " ")

		var chatID int64 = -1002121122014

		for _, word := range words {
			if word != "" {
				if word == "Проблема:" {
					msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("[@%s] "+update.Message.Text, update.Message.Chat.UserName))
					log.Printf("[%d] %s, %s", update.Message.Chat.ID, update.Message.Text, update.Message.Chat.UserName)
					_, err := bot.Send(msg)

					if err != nil {
						fmt.Println(err)
					}
					break
				}
			}

		}
		log.Printf("[%d] %s", update.Message.Chat.ID, update.Message.Text)
		bot.Send(msg)

	}
}
