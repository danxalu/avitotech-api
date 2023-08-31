package cfg

import (
	"github.com/spf13/viper"
)

type (
	Server struct {
		Address string `mapstructure:"address" validate:"required"`
	}

	Postgres struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     string `mapstructure:"port" validate:"required"`
		DBName   string `mapstructure:"name" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
	}

	Config struct {
		Server   Server   `mapstructure:"server" validate:"required"`
		Postgres Postgres `mapstructure:"db" validate:"required"`
	}
)


func ApplyConfig() (*Config, error) {
	viper.SetConfigFile("cfg.yml")

	viper.AddConfigPath("./app")
	viper.AddConfigPath(".")

	viper.SetDefault("postgres.host", "db")
	viper.SetDefault("postgres.port", "5432")
	viper.SetDefault("postgres.name", "postgres")
	viper.SetDefault("postgres.user", "admin")
	viper.SetDefault("postgres.password", "admin")

	viper.SetDefault("server.address", "0.0.0.0:9000")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
