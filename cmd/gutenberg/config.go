package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config struct {
	Server   Server   `json:"server"`
	Database Database `json:"database"`
	Logger   Logger   `json:"logger"`
	Env      string   `json:"env"`
}

type Server struct {
	HTTP string `json:"http"`
}

type Database struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Logger holds configuration required to customize logging for dex.
type Logger struct {
	// Level sets logging level severity.
	Level string `json:"level"`

	// Format specifies the format to be used for logging.
	Format string `json:"format"`
}

func readConfig() (*Config, error) {
	viper.SetDefault("server.http", "0.0.0.0:3000")
	viper.SetDefault("database.name", "gutenberg")
	viper.SetDefault("database.url", "localhost")
	viper.SetDefault("logger.level", "debug")
	viper.SetDefault("logger.format", "text")
	viper.SetDefault("workers.maxWorkers", 4)
	viper.SetDefault("workers.maxQueueSize", 16)
	viper.SetDefault("env", "dev")

	viper.SetEnvPrefix("GUTENBERG")
	viper.BindEnv("env")
	viper.BindEnv("server.http")
	viper.BindEnv("database.url")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	env := viper.GetString("env")
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s/", app))
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", app))
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, err
}
