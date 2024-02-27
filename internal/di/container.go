package di

import (
	"github.com/Xenous-Inc/finapp-telegram-bot/internal/utils/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Container struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Container {
	return &Container{cfg: cfg}
}

func (c *Container) StartBot() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel, error) {
	bot, err := tgbotapi.NewBotAPI(c.cfg.Token)
	if err != nil {
		panic(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, _ := bot.GetUpdatesChan(updateConfig)

	return bot, updates, err
}
