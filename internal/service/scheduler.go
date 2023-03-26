package service

import (
	"github.com/lefes/discord-message-scheduler/internal/repository"
	"github.com/lefes/discord-message-scheduler/pkg/logger"
)

type SchedulerService struct {
	repo repository.Message
}

func NewSchedulerService(repo repository.Message) *SchedulerService {
	return &SchedulerService{repo: repo}
}

func (s *SchedulerService) HelloWorld() error {
	err := s.repo.SetKey("hello", "world")
	if err != nil {
		logger.Error(err)

		return err
	}

	answer, err := s.repo.GetKey("hello")
	if err != nil {
		logger.Error(err)

		return err
	}
	logger.Info(answer)

	err = s.repo.DeleteKey("hello")
	if err != nil {
		logger.Error(err)

		return err
	}

	answer, err = s.repo.GetKey("hello")
	if err != nil {
		logger.Error(err)

		return err
	}

	logger.Info(answer)

	return nil
}
