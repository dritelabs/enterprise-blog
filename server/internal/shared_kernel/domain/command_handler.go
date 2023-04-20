package domain

import "context"

type CommandHandler[C any] interface {
	Execute(ctx context.Context, c C) error
}
