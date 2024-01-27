package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/theghostmac/maclibrarybot/internal/searcher"
)

// HandleSearchCommand initiates a book search.
func HandleSearchCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	session := GetSession(update.Message.From.ID)
	switch session.State {
	case StateNone:
		// Initiate book search
		promptForBook(bot, update.Message.Chat.ID)
		session.State = StateAwaitingBookSearch
	case StateAwaitingBookSearch:
		// Process the book search query
		bookName := update.Message.Text
		books, err := searcher.SearchBooks(bookName)
		if err != nil {
			sendErrorMessage(bot, update.Message.Chat.ID, "Sorry, I couldn't search for books due to an error.")
			session.State = StateNone
			return
		}
		session.LastSearchResults = books
		presentBookOptions(bot, update.Message.Chat.ID, books)
		// No state change here, we're still awaiting book selection
	default:
		// Handle other states or reset
		session.State = StateNone
	}
}
