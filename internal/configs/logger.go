package configs

import (
	"github.com/caarlos0/env/v7"
	"log/slog"
)

var lvls map[string]slog.Leveler = map[string]slog.Leveler{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
}

type LoggerConfig struct {
	Level  slog.Leveler
	level  string `env:"LOG_LEVEL" envDefault:"debug"`
	Format string `env:"LOG_FORMAT" envDefault:"json"`
}

func NewLoggerConfig() *LoggerConfig {

	cfg := LoggerConfig{}

	if err := env.Parse(&cfg); err != nil {
		slog.Error("logger config parse error", slog.Any("err", err))
	}

	cfg.Level = lvls[cfg.level]

	return &cfg
}
