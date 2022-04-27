package ethereum

import (
	"github.com/sirupsen/logrus"
	"nn-blockchain-api/pkg/errors"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go

type Service interface {
}

type service struct {
	log *logrus.Logger
}

func NewService(log *logrus.Logger) (Service, error) {
	if log == nil {
		return nil, errors.NewInternal("invalid logger")
	}
	return &service{log: log}, nil
}
