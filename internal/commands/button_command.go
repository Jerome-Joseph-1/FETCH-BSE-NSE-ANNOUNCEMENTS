package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	RegisterCommand("button", ButtonCommand)
	RegisterComponentHandler("example_button", ButtonHandler)
}

func ButtonCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Content: "Click the button below!",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Click Me",
						Style:    discordgo.PrimaryButton,
						CustomID: "example_button",
					},
				},
			},
		},
	})
}

func ButtonHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "You clicked the button! User ID: " + i.Member.User.ID,
		},
	})
}