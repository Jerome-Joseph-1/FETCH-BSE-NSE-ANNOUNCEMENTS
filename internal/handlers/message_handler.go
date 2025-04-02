package handlers

import (
	"JeromeJoseph/discord-bot/internal/commands"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	prefix := "!"
	if strings.HasPrefix(m.Content, prefix) {
		content := strings.TrimPrefix(m.Content, prefix)
		args := strings.Fields(content)
		if len(args) == 0 {
			return
		}

		commandName := args[0]
		args = args[1:]

		if handler := commands.GetCommand(commandName); handler != nil {
			handler(s, m, args)
		}
	}
}
