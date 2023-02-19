package db

import (
	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// NewStorage where do we store notes
func NewStorage(logger *zerolog.Logger, config *viper.Viper) *redis.Client {
	storage := redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis"),
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	logger.Debug().Msg("Redis start")

	go Daemon(storage, logger, config)

	return storage
}
