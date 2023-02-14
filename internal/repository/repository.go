package repository

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
)

type Storage struct {
	db     *redis.Client
	logger *zerolog.Logger
}

const (
	Day   = time.Hour * 24
	Week  = Day * 7
	Month = Day * 30
	Year  = Day * 365
	Min   = time.Second * 300 //5min
)

func NewRepository(conn *redis.Client, logger *zerolog.Logger) *Storage {
	return &Storage{
		db:     conn,
		logger: logger,
	}
}

// GetAll getting the record
func (s Storage) GetAll(number string) (map[string]string, error) {
	result, err := s.db.HGetAll(number).Result()
	if err != nil {
		s.logger.Error().Str("err", err.Error()).Msg("List note")
		return result, err
	}
	return result, err
}

// GetHMGet getting the field from key
func (s Storage) GetHMGet(key string, field string) ([]interface{}, error) {
	result, err := s.db.HMGet(key, field).Result()
	if err != nil {
		s.logger.Debug().Err(err).Msg("Сannot find user email")
		return nil, err
	}
	return result, err
}

// HSet устанавливаем параметры в хеше
func (s Storage) HSet(key, field string, value interface{}) bool {
	boolState := s.db.HSet(key, field, value)
	return boolState.Val()
}

// HMSet устанавливаем несколько параметров в хеше
func (s Storage) HMSet(key string, fields map[string]interface{}) error {
	err := s.db.HMSet(key, fields).Err()
	return err
}

// Expire установка времени жизни ключа
func (s Storage) Expire(key string, expiration time.Duration) error {
	err := s.db.PExpire(key, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
