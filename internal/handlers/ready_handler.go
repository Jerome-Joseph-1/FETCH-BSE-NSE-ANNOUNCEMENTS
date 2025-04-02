package handlers 

import (
	"JeromeJoseph/discord-bot/internal/logger"
	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	logger.Info.Printf("Bot is ready! Logged in as %s#%s", r.User.Username, r.User.Discriminator)
	s.UpdateGameStatus(0, "!help for commands")
}