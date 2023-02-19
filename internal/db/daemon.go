package db

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// Daemon a daemon that cleans user notes
func Daemon(storage *redis.Client, logger *zerolog.Logger, config *viper.Viper) error {
	clearVacuum := config.GetInt("VacuumClearSec")
	timeout := time.Duration(clearVacuum) * time.Second
	for {
		<-time.After(timeout)
		logger.Debug().Str("timeout", timeout.String()).Msg("Daemon check ttl init")
		userhash, err := storage.HGetAll("usersm").Result()
		if err != nil {
			logger.Fatal().Msg("Daemon check get all")
		}

		for hash := range userhash {
			result, err := storage.HGetAll(hash).Result()
			if err != nil {
				logger.Error().Str("error", err.Error()).Msg("Daemon check get user hash")
			}
			if len(result) > 0 {
				logger.Debug().Msg("Daemon get user notes")
				for note := range result {
					logger.Debug().Msg("Daemon looking for note in storage")
					notes, err := storage.HGetAll(note).Result()
					if err != nil {
						logger.Error().Str("error", err.Error()).Msg("Daemon not found note in storage")
					}
					if len(notes) == 0 {
						logger.Debug().Msg("Daemon found user note, need delete her")
						_, err := storage.HDel(hash, note).Result()
						if err != nil {
							logger.Error().Str("error", err.Error()).Msg("Daemon not remove keys in usersm")
						}
						logger.Info().Str("note", note).
							Str("user", hash).
							Msg("Daemon remove not live note")
					}

				}
			}
		}
		logger.Debug().Msg("Daemon check ttl end")
	}
}
