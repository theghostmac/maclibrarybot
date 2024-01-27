package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/theghostmac/maclibrarybot/internal/searcher"
)

func HandleBookSelection(bot *tgbotapi.BotAPI, update tgbotapi.Update, selection int) {
	session := GetSession(update.Message.From.ID)

	if session.State == StateAwaitingBookSearch && selection >= 1 && selection <= len(session.LastSearchResults) {
		// User has made a valid selection, so provide details for the selected book.
		selectedBook := session.LastSearchResults[selection-1]
		sendBookDetails(bot, update.Message.Chat.ID, selectedBook)
		session.State = StateNone // Reset the state after handling the selection
	} else {
		sendErrorMessage(bot, update.Message.Chat.ID, "Please select a valid book number.")
	}
}

// sendBookDetails sends the details of the selected book to the user.
func sendBookDetails(bot *tgbotapi.BotAPI, chatID int64, book searcher.Book) {
	messageText := fmt.Sprintf("Title: %s\nAuthor: %s\nLink: %s", book.Title, book.Author, book.Link)
	msg := tgbotapi.NewMessage(chatID, messageText)
	bot.Send(msg)
}
