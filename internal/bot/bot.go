package bot

import (
	"JeromeJoseph/discord-bot/config"
	"JeromeJoseph/discord-bot/internal/handlers"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session
	Config  *config.Config
}

func New(cfg *config.Config) (*Bot, error) {
	session, err := discordgo.New("Bot " + cfg.Bot.Token)

	if err != nil {
		return nil, err
	}

	return &Bot{
		Session: session,
		Config:  cfg,
	}, nil
}

func (b *Bot) Start() error {
	b.Session.AddHandler(handlers.Ready)

	return b.Session.Open()
}

func (b *Bot) Stop() error {
	return b.Session.Close()
}
