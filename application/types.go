package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-youtube/services/youtube"
)

type Application interface {
	Run() error
	Shutdown()
}

type Impl struct {
	youtubeService youtube.Service
	broker         amqp.MessageBroker
}
