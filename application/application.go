package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-youtube/models/constants"
	"github.com/kaellybot/kaelly-youtube/repositories/videasts"
	"github.com/kaellybot/kaelly-youtube/services/youtube"
	"github.com/kaellybot/kaelly-youtube/utils/databases"
	"github.com/kaellybot/kaelly-youtube/utils/insights"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Impl, error) {
	// misc
	broker := amqp.New(constants.RabbitMQClientID, viper.GetString(constants.RabbitMQAddress))
	db := databases.New()
	probes := insights.NewProbes(broker.IsConnected, db.IsConnected)
	prom := insights.NewPrometheusMetrics()

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
		db:             db,
		probes:         probes,
		prom:           prom,
	}, nil
}

func (app *Impl) Run() error {
	app.probes.ListenAndServe()
	app.prom.ListenAndServe()

	if err := app.db.Run(); err != nil {
		return err
	}

	if err := app.broker.Run(); err != nil {
		return err
	}

	return app.youtubeService.Consume()
}

func (app *Impl) Shutdown() {
	app.broker.Shutdown()
	app.db.Shutdown()
	app.prom.Shutdown()
	app.probes.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
