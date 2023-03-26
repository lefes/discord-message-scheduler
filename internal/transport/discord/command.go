package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lefes/discord-message-scheduler/pkg/logger"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Ping the bot",
			Type:        discordgo.ChatApplicationCommand,
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping": PingCommand,
	}

	registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))
)

func PingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
	if err != nil {
		logger.Error(err)
	}

}
