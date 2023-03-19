package service

import (
	"sql-service/domain"
	"sql-service/model"
	"sql-service/utils"

	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

type Service struct {
	Storage domain.IStorage
	Log     *zap.Logger
}

type IService interface {
	Console(request model.Request) model.Response
}

func (s *Service) Console(request model.Request) model.Response {
	var response model.Response

	const (
		DML = "DML"
	)

	defaulttypes := []string{"DQL", "TCL", "DDL", "DQL", "PLUGIN"}

	if slices.Contains(defaulttypes, utils.SQLType(request.SQL)) {
		response = s.Storage.Console(request.SQL)
	}

	if DML == utils.SQLType(request.SQL) {
		response = s.Storage.ConsoleDML(request.SQL)
	}

	return response
}
