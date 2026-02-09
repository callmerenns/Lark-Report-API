package service

import (
	"context"
	"errors"

	"github.com/tsaqif-19/lark-report-api/internal/domain"
	"github.com/tsaqif-19/lark-report-api/internal/logger"
	"github.com/tsaqif-19/lark-report-api/internal/repository"
	"go.uber.org/zap"
)

type RecordService struct {
	repo repository.RecordRepository
}

func NewRecordService(r repository.RecordRepository) *RecordService {
	return &RecordService{repo: r}
}

func (s *RecordService) CreateRecord(
	ctx context.Context,
	rec domain.Record,
) (int64, error) {

	id, err := s.repo.Create(ctx, rec)
	if err != nil {

		if errors.Is(err, context.DeadlineExceeded) {
			logger.Log.Error.Warn(
				"record_create_timeout",
				zap.String("service", "record"),
			)
		} else {
			logger.Log.Error.Error(
				"record_create_failed",
				zap.Error(err),
				zap.String("service", "record"),
			)
		}

		return 0, err
	}

	return id, nil
}
