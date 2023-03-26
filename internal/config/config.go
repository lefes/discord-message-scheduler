package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"github.com/lefes/discord-message-scheduler/pkg/logger"
)

func Init() (*Config, error) {
	err := loadDotEnv()
	if err != nil {
		return nil, err
	}

	var cfg Config

	cfg.populateDefaults()

	err = cfg.setFromEnv()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) setFromEnv() error {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("discord")
	if viper.GetString("token") == "" {
		return fmt.Errorf("discord token is not set")
	}

	cfg.Discord = DiscordConfig{
		Token:        viper.GetString("token"),
		ReadMessages: viper.GetBool("read_messages"),
		GuildID:      viper.GetString("guild_id"),
	}

	viper.SetEnvPrefix("db")
	cfg.DB = DBConfig{
		Path: viper.GetString("path"),
	}

	viper.SetEnvPrefix("http")
	cfg.HTTP = HTTPConfig{
		Host:              viper.GetString("host"),
		Port:              viper.GetString("port"),
		ReadHeaderTimeout: viper.GetInt("read_header_timeout"),
	}

	return nil
}

func (cfg *Config) populateDefaults() {
	// DB
	viper.SetEnvPrefix("db")
	viper.SetDefault("path", DefaultDBPath)

	// HTTP
	viper.SetEnvPrefix("http")
	viper.SetDefault("host", DefaultHTTPHost)
	viper.SetDefault("port", DefaultHTTPPort)
	viper.SetDefault("read_header_timeout", DefaultReadHeaderTimeout)

	// Discord
	viper.SetEnvPrefix("discord")
	viper.SetDefault("read_messages", DefaultDiscordReadMessages)
	viper.SetDefault("guild_id", DefaultDiscordGuildID)
}

func loadDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		logger.Warn("no .env file found")
	}

	return nil
}
