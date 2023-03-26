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

func (c *Client) Start() error {

	// Register handlers
	c.discordClient.AddHandler(handlerReady)
	c.discordClient.AddHandler(handlerInteractionCreate)

	err := c.discordClient.Open()
	if err != nil {
		return err
	}

	// Register commands
	for i, v := range commands {
		cmd, err := c.discordClient.ApplicationCommandCreate(c.discordClient.State.User.ID, c.cfg.GuildID, v)
		if err != nil {
			return fmt.Errorf("cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd

	}

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
