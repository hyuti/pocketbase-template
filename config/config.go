package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		Log   `yaml:"logger"`
		DB    `yaml:"db"`
		Redis `yaml:"redis"`
	}

	// App -.
	App struct {
		Name     string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version  string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		Debug    bool   `env-required:"true" yaml:"debug" env:"APP_DEBUG"`
		Email    string `env-required:"true" yaml:"email" env:"APP_EMAIL"`
		Password string `env-required:"true" yaml:"password" env:"APP_PASSWORD"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// DB -.
	DB struct {
		DataDir string `env-required:"true" yaml:"dataDir" env:"DB_DATA_DIR"`
	}

	// Redis -.
	Redis struct {
		URL      string `env-required:"true"                 env:"REDIS_DB_URL"`
		Password string `env-required:"true"                 env:"REDIS_DB_PASSWORD"`
		DB       int    `env-required:"true"                 env:"REDIS_DB"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"logLevel" env:"LOG_LEVEL"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
