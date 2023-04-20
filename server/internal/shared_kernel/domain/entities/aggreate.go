package entities

import "github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"

type Aggregate struct {
	events []domain.Event
}

func (a *Aggregate) Apply(e domain.Event) {
	a.events = append(a.events, e)
}

func (a *Aggregate) Events() []domain.Event {
	return a.events
}
