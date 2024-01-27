package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/theghostmac/maclibrarybot/internal/searcher"
	"go.uber.org/zap"
)

// sendBotMessage is a helper function to send a message using the bot.
func sendBotMessage(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	_, err := bot.Send(msg)
	if err != nil {
		logger.Error("Failed to send message", zap.Error(err))
	}
}

// promptForBook sends a message to the user asking them to enter a book title.
func promptForBook(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "What book are you looking for?")
	bot.Send(msg)
}

// presentBookOptions formats the found books as a list and sends it to the user.
func presentBookOptions(bot *tgbotapi.BotAPI, chatID int64, books []searcher.Book) {
	messageText := "I found these books:\n\n"
	for i, book := range books {
		messageText += fmt.Sprintf("%d. %s by %s\n", i+1, book.Title, book.Author)
	}
	messageText += "\nSelect a book by typing its number."
	msg := tgbotapi.NewMessage(chatID, messageText)
	bot.Send(msg)
}

// sendErrorMessage is a helper function to send an error message to the user.
func sendErrorMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
}
