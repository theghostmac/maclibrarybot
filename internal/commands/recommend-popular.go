package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/theghostmac/maclibrarybot/internal/recommender"
)

func HandleRecommendPopular(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	recommender.SendPopularBooks(bot, update.Message.Chat.ID)
}
