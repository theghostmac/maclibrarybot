package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

var logger, _ = zap.NewDevelopment()

var startMessage = "👋 Hello!\n\n📚 If you have an interest in a book, just write its name and I will try to find it for you.\n✨ I can also recommend you the most popular and exciting books!\n\n❗️ This bot does not violate copyright, it will simply find for you any book from open Internet sources!"

func HandleStartCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Create a new ReplyKeyboardMarkup struct
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🔍 Search A Book"),
			tgbotapi.NewKeyboardButton("🔥 Popular Books"),
			// Add more buttons here if needed.
		),
	)

	keyboard.ResizeKeyboard = true // This optimizes the keyboard size for the screen

	// Create a new message with the ReplyKeyboardMarkup
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, startMessage)
	msg.ReplyMarkup = keyboard

	// Send the message with the custom keyboard
	sendBotMessage(bot, msg)
}
