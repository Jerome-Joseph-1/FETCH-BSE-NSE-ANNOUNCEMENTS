package main

import (
	"JeromeJoseph/discord-bot/config"
	"JeromeJoseph/discord-bot/internal/bot"
	_ "JeromeJoseph/discord-bot/internal/commands"
	"JeromeJoseph/discord-bot/internal/logger"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configPath := flag.String("config", "config.yaml", "config.yaml")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	logger.Init(cfg.Logging.Level)

	b, err := bot.New(cfg)
	if err != nil {
		logger.Error.Fatalf("Failed to create bot: %v", err)
	}

	if err := b.Start(cfg); err != nil {
		logger.Error.Fatalf("Failed to start bot: %v", err)
	}

	logger.Info.Println("Bot is now running. Press CTRL+C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	logger.Info.Println("Shutting down...")
	b.Stop()
}
