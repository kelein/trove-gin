package service

import (
	"github.com/kelein/trove-gin/internal/repository"
	"github.com/kelein/trove-gin/pkg/jwt"
	"github.com/kelein/trove-gin/pkg/log"
	"github.com/kelein/trove-gin/pkg/sid"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}
