package main

import (
	"github.com/Xenous-Inc/finapp-telegram-bot/internal/di"
	"github.com/Xenous-Inc/finapp-telegram-bot/internal/service"
	"github.com/Xenous-Inc/finapp-telegram-bot/internal/utils/config"
	"github.com/Xenous-Inc/finapp-telegram-bot/internal/utils/flags"

	"github.com/4kayDev/logger"
	"github.com/4kayDev/logger/log"
)

func main() {
	flags := flags.MustParseFlags()
	cfg := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	logger.ConfigureZeroLogger()
	container := di.New(cfg)

	bot, updates, err := container.StartBot()
	if err != nil {
		panic(err)
	}

	log.Debug("Registring service")

	service.SendMsg(bot, updates)
}
