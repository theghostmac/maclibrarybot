package recommender

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func SendPopularBooks(bot *tgbotapi.BotAPI, chatID int64) {
	// Mock-up function, you'll need to replace this with a real implementation.
	messageText := "Here are some popular books:\n1. Book One\n2. Book Two\n3. Book Three\n"
	msg := tgbotapi.NewMessage(chatID, messageText)
	bot.Send(msg)
}
