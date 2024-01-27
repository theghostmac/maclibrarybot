package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Check for a message update
	if update.Message != nil {
		// Retrieve the user session
		session := GetSession(update.Message.From.ID)

		switch session.State {
		case StateAwaitingBookSearch:
			// User is expected to enter the name of a book
			if !update.Message.IsCommand() {
				HandleSearchCommand(bot, update)
				return
			}
		case StateNone:
			// Handle other commands like /start
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					HandleStartCommand(bot, update)
					// Add more commands as needed
				}
			}
		}

		// If the user is expected to select a book by number, handle that
		if selection, err := strconv.Atoi(update.Message.Text); err == nil {
			// The user has replied with a number, handle book selection
			HandleBookSelection(bot, update, selection)
		}
	}
}
