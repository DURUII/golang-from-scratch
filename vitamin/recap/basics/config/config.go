package config

import (
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
)

var (
	cfg *viper.Viper
)

func init() {
	cfg = viper.New()
	cfg.SetConfigName("config")
	cfg.AddConfigPath(".")
	cfg.SetConfigType("yaml")
	if err := cfg.ReadInConfig(); err != nil {
		panic(err)
	}
	slog.Info("config loaded", "files", cfg.ConfigFileUsed())
}

func GetString(key string) string {
	return cfg.GetString(key)
}

func GetInt(key string) int {
	return cfg.GetInt(key)
}
