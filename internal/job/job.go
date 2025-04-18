package job

import (
	"github.com/kelein/trove-gin/internal/repository"
	"github.com/kelein/trove-gin/pkg/jwt"
	"github.com/kelein/trove-gin/pkg/log"
	"github.com/kelein/trove-gin/pkg/sid"
)

type Job struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewJob(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
) *Job {
	return &Job{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
