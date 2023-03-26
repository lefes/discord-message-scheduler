package config

type (
	Config struct {
		Discord DiscordConfig
		HTTP    HTTPConfig
		DB      DBConfig
		Log     LogConfig
	}

	DiscordConfig struct {
		Token        string
		ReadMessages bool
		GuildID      string
	}

	HTTPConfig struct {
		Host              string
		Port              string
		ReadHeaderTimeout int
	}

	DBConfig struct {
		Path string
	}

	LogConfig struct {
		Level  string
		Format string
	}
)
