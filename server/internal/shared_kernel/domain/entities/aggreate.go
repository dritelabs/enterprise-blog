package entities

import (
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type Aggregate struct {
	events   []domain.Event
	eventBus domain.EventBus
}

func (a *Aggregate) Apply(e domain.Event) {
	a.events = append(a.events, e)
}

func (a *Aggregate) Commit() {
	for _, ev := range a.events {
		a.eventBus.Publish(ev.String(), ev)
	}

	a.events = []domain.Event{}
}

func (a *Aggregate) Events() []domain.Event {
	return a.events
}

func (a *Aggregate) WithEventBus(eventBus domain.EventBus) {
	a.eventBus = eventBus
}
