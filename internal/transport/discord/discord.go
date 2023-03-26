package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/lefes/discord-message-scheduler/internal/config"
)

type Client struct {
	discordClient *discordgo.Session
	cfg           *config.DiscordConfig
}

func NewClient(cfg *config.DiscordConfig) (*Client, error) {
	discordClient, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return nil, err
	}

	if cfg.ReadMessages {
		discordClient.Identify.Intents = discordgo.IntentsGuildMessages
	}

	return &Client{
		discordClient: discordClient,
		cfg:           cfg,
	}, nil
}

func (c *Client) Start(h Handler, com Command) error {

	// Register handlers
	h.RegisterHandlers(c.discordClient)

	err := c.discordClient.Open()
	if err != nil {
		return err
	}

	// Register commands
	// TODO IMIDIATELY: fix access to regitster commands
	com.RegisterCommands(c.discordClient)

	return nil
}

func (c *Client) Shutdown() error {
	for _, v := range registeredCommands {
		err := c.discordClient.ApplicationCommandDelete(c.discordClient.State.User.ID, c.cfg.GuildID, v.ID)
		if err != nil {
			return fmt.Errorf("cannot delete '%v' command: %v", v.Name, err)
		}
	}
	return c.discordClient.Close()
}
