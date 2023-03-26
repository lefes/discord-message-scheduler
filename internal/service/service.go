package service

import "github.com/lefes/discord-message-scheduler/internal/repository"

type Scheduler interface {
	HelloWorld() error
}

type Services struct {
	Scheduler Scheduler
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		Scheduler: NewSchedulerService(deps.Repos.Message),
	}
}
