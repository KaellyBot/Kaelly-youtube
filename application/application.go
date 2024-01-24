package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-youtube/models/constants"
	"github.com/kaellybot/kaelly-youtube/repositories/videasts"
	"github.com/kaellybot/kaelly-youtube/services/youtube"
	"github.com/kaellybot/kaelly-youtube/utils/databases"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Impl, error) {
	// misc
	db, err := databases.New()
	if err != nil {
		return nil, err
	}

	broker, err := amqp.New(constants.RabbitMQClientID, viper.GetString(constants.RabbitMQAddress), nil)
	if err != nil {
		return nil, err
	}

	// repositories
	videastRepo := videasts.New(db)

	// services
	youtubeService, err := youtube.New(videastRepo, broker)
	if err != nil {
		return nil, err
	}

	return &Impl{
		youtubeService: youtubeService,
		broker:         broker,
	}, nil
}

func (app *Impl) Run() error {
	return app.youtubeService.Consume()
}

func (app *Impl) Shutdown() {
	app.broker.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
