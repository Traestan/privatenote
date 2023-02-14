package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

// NewLogger returns new instance *zerolog.Logger with default settings
func NewLogger(config *viper.Viper) *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano}
	lvl, err := zerolog.ParseLevel(config.GetString("logLevel"))
	if err != nil {
		panic(err)
	}
	log := zerolog.New(output).Level(lvl).With().Timestamp().Logger()
	return &log
}
