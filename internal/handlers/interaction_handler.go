package handlers

import (
	"JeromeJoseph/discord-bot/internal/commands"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h := commands.GetSlashCommand(i.ApplicationCommandData().Name); h != nil {
			h(s, i)
		}
	case discordgo.InteractionMessageComponent:
		if h := commands.GetComponentHandler(i.MessageComponentData().CustomID); h != nil {
			h(s, i)
		}
	}
}
