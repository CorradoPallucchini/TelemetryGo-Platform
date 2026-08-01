package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CorradoPallucchini/telemetrygo-platform/internal/domain"
)

var ErrStationNotFound = errors.New("station not found")

type StationRepository struct {
	db *sql.DB
}

func NewStationRepository(db *sql.DB) *StationRepository {
	return &StationRepository{db: db}
}

func (r *StationRepository) Create(ctx context.Context, station *domain.Station) error {
	query := `
		INSERT INTO stations (id, name, location, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(ctx, query, station.ID, station.Name, station.Location, station.CreatedAt, station.UpdatedAt)
	return err
}

func (r *StationRepository) GetByID(ctx context.Context, id string) (*domain.Station, error) {
	query := `
		SELECT id, name, location, created_at, updated_at
		FROM stations
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var s domain.Station
	err := row.Scan(&s.ID, &s.Name, &s.Location, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Ritorniamo il nostro errore personalizzato se il record non esiste
			return nil, ErrStationNotFound
		}
		return nil, err
	}
	return &s, nil
}

func (r *StationRepository) List(ctx context.Context) ([]*domain.Station, error) {
	query := `SELECT id, name, location, created_at, updated_at FROM stations`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stations []*domain.Station

	for rows.Next() {
		var s domain.Station
		if err := rows.Scan(&s.ID, &s.Name, &s.Location, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		stations = append(stations, &s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stations, nil
}

func (r *StationRepository) Update(ctx context.Context, station *domain.Station) error {
	query := `
		UPDATE stations
		SET name = $1, location = $2, updated_at = $3
		WHERE id = $4
	`
	res, err := r.db.ExecContext(ctx, query, station.Name, station.Location, station.UpdatedAt, station.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrStationNotFound
	}

	return nil
}

func (r *StationRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM stations WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
