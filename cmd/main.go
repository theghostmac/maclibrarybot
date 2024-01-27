package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/theghostmac/maclibrarybot/config"
	"github.com/theghostmac/maclibrarybot/internal/handlebot"
	"go.uber.org/zap"
)

var logger, _ = zap.NewDevelopment()

func main() {
	// Initialize environment variables.
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Could not load config: ", zap.Error(err))
	}

	// Initialize the telegram bot.
	bot, err := tgbotapi.NewBotAPI(cfg.MacLibraryBotToken)
	if err != nil {
		logger.Fatal("Error creating Telegram bot: ", zap.Error(err))
	}
	logger.Info("Authorized on account", zap.String("bot", bot.Self.UserName))

	// Telegram updates loop.
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			// Ignore any non-Message updates.
			continue
		}

		// Pass the update to the handler.
		handlebot.HandleUpdate(bot, update)

		// Logging the message for debugging purposes.
		logger.Info("Received a message", zap.String("from", update.Message.From.UserName), zap.String("text", update.Message.Text))
	}
}
