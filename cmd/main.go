package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"gitlab.com/privatenote/api"
	"gitlab.com/privatenote/internal/config"
	"gitlab.com/privatenote/internal/db"
	"gitlab.com/privatenote/internal/encode"
	"gitlab.com/privatenote/internal/logger"
	"gitlab.com/privatenote/internal/model"
	"gitlab.com/privatenote/internal/repository"
	"go.uber.org/dig"
)

func main() {
	contnr := containerBuild()
	err := contnr.Invoke(func(
		logger *zerolog.Logger,
		config *viper.Viper,
		server *api.Server,
	) {
		logger.Info().Msg("Prepare to server start")
		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Fatal().Err(err).Msg("error starting server")
			}
		}()
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()
		logger.Info().Msg("Prepare to graceful shutdown")
		err := server.Shutdown(ctx)
		if err != nil {
			logger.Fatal().Err(err).Msg("error shutdown server")
		}
		logger.Info().Msg("Bye!")
		os.Exit(0)
	})

	if err != nil {
		panic(err)
	}
}
func containerBuild() (container *dig.Container) {
	container = dig.New()
	if err := container.Provide(config.NewConfig); err != nil {
		panic(err)
	}
	if err := container.Provide(logger.NewLogger); err != nil {
		panic(err)
	}
	if err := container.Provide(db.NewStorage); err != nil {
		panic(err)
	}
	if err := container.Provide(repository.NewRepository); err != nil {
		panic(err)
	}
	if err := container.Provide(api.NewRouter); err != nil {
		panic(err)
	}
	if err := container.Provide(encode.NewURLEncoder); err != nil {
		panic(err)
	}
	if err := container.Provide(model.NewFactory); err != nil {
		panic(err)
	}
	if err := container.Provide(api.NewServer); err != nil {
		panic(err)
	}
	return container
}
