package domain

import (
	"context"
	"time"
)

type Station struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StationRepository interface {
	Create(ctx context.Context, station *Station) error

	GetByID(ctx context.Context, id string) (*Station, error)

	List(ctx context.Context) ([]*Station, error)

	Update(ctx context.Context, station *Station) error

	Delete(ctx context.Context, id string) error
}
