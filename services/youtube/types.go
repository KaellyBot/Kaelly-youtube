package youtube

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-youtube/repositories/videasts"
)

const (
	Routingkey = "news.youtube"
)

type Service interface {
	Consume() error
}

type Impl struct {
	videastRepo videasts.Repository
	broker      amqp.MessageBroker
}
