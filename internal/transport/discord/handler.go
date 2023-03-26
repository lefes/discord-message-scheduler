package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lefes/discord-message-scheduler/pkg/logger"
)

type Handler struct {
	commands *Command
}

func NewHandlers(commands *Command) *Handler {
	return &Handler{
		commands: commands,
	}
}

func (h *Handler) RegisterHandlers(s *discordgo.Session) {
	s.AddHandler(h.handlerReady)
	s.AddHandler(h.handlerInteractionCreate)
}

func (h *Handler) handlerReady(s *discordgo.Session, r *discordgo.Ready) {
	logger.Info("Bot is up!")
	logger.Infof("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
}

func (h *Handler) handlerInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping":  h.commands.PingCommand,
		"hello": h.commands.HelloWorld,
	}
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}
