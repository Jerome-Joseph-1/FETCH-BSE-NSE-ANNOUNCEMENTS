package commands 

import (
	"github.com/bwmarrin/discordgo"
)

type MessageHandler func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
type InteractionHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)

var textCommands = make(map[string]MessageHandler)
var slashCommands = make(map[string]InteractionHandler)
var componentHandlers = make(map[string]InteractionHandler)

func RegisterCommand(name string, handler MessageHandler) {
	textCommands[name] = handler
}

func RegisterSlashCommand(name string, handler InteractionHandler) {
	slashCommands[name] = handler
}

func RegisterComponentHandler(customID string, handler InteractionHandler) {
	componentHandlers[customID] = handler
}

func GetCommand(name string) MessageHandler {
	return textCommands[name]
}

func GetSlashCommand(name string) InteractionHandler {
	return slashCommands[name]
}

func GetComponentHandler(customID string) InteractionHandler {
	return componentHandlers[customID]
}