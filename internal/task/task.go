package task

import (
	"github.com/kelein/trove-gin/internal/repository"
	"github.com/kelein/trove-gin/pkg/jwt"
	"github.com/kelein/trove-gin/pkg/log"
	"github.com/kelein/trove-gin/pkg/sid"
)

type Task struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewTask(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
) *Task {
	return &Task{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
