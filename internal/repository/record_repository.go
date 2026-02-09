package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tsaqif-19/lark-report-api/internal/domain"
	"github.com/tsaqif-19/lark-report-api/internal/logger"
	"go.uber.org/zap"
)

type RecordRepository interface {
	Create(ctx context.Context, record domain.Record) (int64, error)
}

type recordRepo struct {
	db *pgxpool.Pool
}

func NewRecordRepository(db *pgxpool.Pool) RecordRepository {
	return &recordRepo{db: db}
}

func (r *recordRepo) Create(
	ctx context.Context,
	rec domain.Record,
) (int64, error) {

	var id int64

	err := r.db.QueryRow(ctx, `
		INSERT INTO records (
			incident_description, machine_model,
			lightbox_type, set_type, types_of_spareparts,
			how_to_handle, collect_dataset, processing_model,
			date_of_incident, time_of_incident, day_of_incident,
			process_reason, created_by, modified_by, status
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15
		)
		RETURNING id
	`,
		rec.IncidentDescription,
		rec.MachineModel,
		rec.LightboxType,
		rec.SetType,
		rec.TypesOfSpareparts,
		rec.HowToHandle,
		rec.CollectDataset,
		rec.ProcessingModel,
		rec.DateOfIncident,
		rec.TimeOfIncident,
		rec.DayOfIncident,
		rec.ProcessReason,
		rec.CreatedBy,
		rec.ModifiedBy,
		rec.Status,
	).Scan(&id)

	if err != nil {

		// ⏱ Context timeout / cancel
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
			logger.Log.Error.Warn(
				"db_insert_record_timeout",
				zap.String("repository", "record"),
			)
		} else if errors.Is(err, pgx.ErrNoRows) {

			logger.Log.Error.Warn(
				"db_insert_returned_no_rows",
				zap.String("repository", "record"),
			)

		} else {

			logger.Log.Error.Error(
				"db_insert_record_failed",
				zap.Error(err),
				zap.String("repository", "record"),
			)
		}

		return 0, err
	}

	return id, nil
}
