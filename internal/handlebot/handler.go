package handlebot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/theghostmac/maclibrarybot/internal/commands"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return // No message to handle.
	}

	switch update.Message.Command() {
	case "start":
		commands.HandleStartCommand(bot, update)
	}
}
