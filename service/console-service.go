package service

import (
	"sql-service/domain"

	"go.uber.org/zap"
)

type Service struct {
	Storage domain.IStorage
	Log     *zap.Logger
}

type IService interface {
	Console()
}

func (s *Service) Console() {
	s.Storage.Console()
}
