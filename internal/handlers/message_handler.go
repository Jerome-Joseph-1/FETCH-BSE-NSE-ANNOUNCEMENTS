package handlers

import (
	"JeromeJoseph/discord-bot/config"
	"JeromeJoseph/discord-bot/internal/logger"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ChannelMonitorHandler(cfg *config.Config) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.ChannelID != cfg.Bot.Channel_ID {
			return
		}

		logger.Info.Printf("Message received in channel: %s", m.Content)

		processMessage(s, m)
	}
}

func processMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!help") {
		s.ChannelMessageSend(m.ChannelID, "YOU ARE BEING HELPED")
	}
}
