package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-youtube/services/youtube"
	"github.com/kaellybot/kaelly-youtube/utils/databases"
	"github.com/kaellybot/kaelly-youtube/utils/insights"
)

type Application interface {
	Run() error
	Shutdown()
}

type Impl struct {
	youtubeService youtube.Service
	broker         amqp.MessageBroker
	db             databases.MySQLConnection
	probes         insights.Probes
	prom           insights.PrometheusMetrics
}
