package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lefes/discord-message-scheduler/pkg/logger"
)

func handlerReady(s *discordgo.Session, r *discordgo.Ready) {
	logger.Info("Bot is up!")
	logger.Infof("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
}

func handlerInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}
