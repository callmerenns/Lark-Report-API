package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tsaqif-19/lark-report-api/internal/domain"
)

type RecordRepository interface {
	Create(ctx context.Context, record domain.Record) (string, error)
}

type recordRepo struct {
	db *pgxpool.Pool
}

func NewRecordRepository(db *pgxpool.Pool) RecordRepository {
	return &recordRepo{db: db}
}

func (r *recordRepo) Create(ctx context.Context, rec domain.Record) (string, error) {
	var id string

	err := r.db.QueryRow(ctx, `
		INSERT INTO records (
			incident_title, incident_description, machine_model,
			lightbox_type, set_type, types_of_spareparts,
			how_to_handle, collect_dataset, processing_model,
			date_of_incident, time_of_incident, day_of_incident,
			process_reason, created_by, modified_by, status
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16
		)
		RETURNING id
	`,
		rec.IncidentTitle,
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

	return id, err
}
