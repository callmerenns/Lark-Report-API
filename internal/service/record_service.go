package service

import (
	"context"

	"github.com/tsaqif-19/lark-report-api/internal/domain"
	"github.com/tsaqif-19/lark-report-api/internal/repository"
)

type RecordService struct {
	repo repository.RecordRepository
}

func NewRecordService(r repository.RecordRepository) *RecordService {
	return &RecordService{repo: r}
}

func (s *RecordService) CreateRecord(ctx context.Context, rec domain.Record) (string, error) {
	return s.repo.Create(ctx, rec)
}
