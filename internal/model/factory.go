package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
	"github.com/traestan/privatenote/internal/encode"
	"github.com/traestan/privatenote/internal/repository"
)

// modelData
type modelData struct {
	storage    *redis.Client
	repository *repository.Storage
	logger     *zerolog.Logger
	encoder    encode.Encoder
}

type Factory struct {
	modelData
}

func NewFactory(storage *redis.Client, repository *repository.Storage, logger *zerolog.Logger, encoder encode.Encoder) *Factory {
	logger.Debug().Msg("Factory start")
	return &Factory{
		modelData: modelData{storage: storage, repository: repository, logger: logger, encoder: encoder},
	}
}

func (m *Factory) NewUser(user *User) *User {
	return newUser(m.modelData, user)
}
func (m *Factory) NewNote(note *Note) *Note {
	return newNote(m.modelData, note)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
