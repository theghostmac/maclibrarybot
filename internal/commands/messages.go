package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

// sendBotMessage is a helper function to send a message using the bot.
func sendBotMessage(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	_, err := bot.Send(msg)
	if err != nil {
		// Log the error. You can replace this with your preferred logging method.
		// For example, if you use zap logger, it would be something like:
		logger.Error("Failed to send message", zap.Error(err))
	}
}
