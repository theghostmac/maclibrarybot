package handlebot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/theghostmac/maclibrarybot/internal/commands"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return // No message to handle.
	}

	// Check if the message text matches the labels on the custom keyboard buttons.
	switch update.Message.Text {
	case "🔍 Search A Book":
		commands.HandleSearchCommand(bot, update)
	case "🔥 Popular Books":
		commands.HandleRecommendPopular(bot, update)
	default:

		switch update.Message.Command() {
		case "start":
			commands.HandleStartCommand(bot, update)
		case "search":
			commands.HandleSearchCommand(bot, update)
		case "recommend":
			commands.HandleRecommendPopular(bot, update)
		}
	}
}
