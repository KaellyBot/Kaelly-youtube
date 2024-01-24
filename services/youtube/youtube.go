package youtube

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-youtube/repositories/videasts"
)

func New(videastRepo videasts.Repository, broker amqp.MessageBroker) (*Impl, error) {
	return &Impl{
		broker:      broker,
		videastRepo: videastRepo,
	}, nil
}

func (service *Impl) Consume() error {
	// TODO
	return service.dispatchYoutubeEvent()
}

func (service *Impl) dispatchYoutubeEvent() error {
	// TODO
	return nil
}
