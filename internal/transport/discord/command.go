package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lefes/discord-message-scheduler/internal/service"
	"github.com/lefes/discord-message-scheduler/pkg/logger"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Ping the bot",
			Type:        discordgo.ChatApplicationCommand,
		},
		{
			Name:        "hello",
			Description: "Hello World",
			Type:        discordgo.ChatApplicationCommand,
		},
	}

	registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))
)

type Command struct {
	service *service.Services
}

func NewCommands(service *service.Services) *Command {
	return &Command{service: service}
}

func (c *Command) RegisterCommands(s *discordgo.Session) {
	for i, command := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, s.State.User.ID, command)
		if err != nil {
			logger.Error(err)
		}
		registeredCommands[i] = cmd
	}
}

func (c *Command) UnregisterCommands(s *discordgo.Session) {
	for _, command := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, s.State.User.ID, command.ID)
		if err != nil {
			logger.Error(err)
		}
	}
}

func (c *Command) PingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

func (c *Command) HelloWorld(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := c.service.Scheduler.HelloWorld()
	if err != nil {
		logger.Error(err)
	}

	logger.Info("Hello World service executed!")
}
