package domain

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/infrastructure/event_stores/models"
)

type EventStore interface {
	List(ctx context.Context) ([]models.Event, error)
	Store(ctx context.Context, e Event) (*models.Event, error)
}
