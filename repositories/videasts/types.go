package videasts

import (
	"github.com/kaellybot/kaelly-youtube/models/entities"
	"github.com/kaellybot/kaelly-youtube/utils/databases"
)

type Repository interface {
	GetVideasts() ([]entities.Videast, error)
}

type Impl struct {
	db databases.MySQLConnection
}
