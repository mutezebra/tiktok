package interaction

import (
	"github.com/Mutezebra/tiktok/app/interaction/domain/repository"
	"github.com/Mutezebra/tiktok/pkg/snowflake"
)

type Service struct {
	Repo repository.InteractionRepository
}

func NewService(service *Service) *Service {
	if service.Repo == nil {
		panic("interaction service.repo should not be nil")
	}
	return service
}

func (srv *Service) GenerateID() int64 {
	return snowflake.GenerateID()
}
