package eventbus

import (
	evbus "github.com/asaskevich/EventBus"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

// uncomment this if a singleton is needed
// var localEventBus *LocalEventBus

type LocalEventBus struct {
	bus evbus.Bus
}

func (e *LocalEventBus) Publish(t string, p domain.Event) {
	e.bus.Publish(t, p)
}

func (e *LocalEventBus) Subscribe(t string, fn domain.EventBusHandler) error {
	return e.bus.Subscribe(t, fn)
}

func (e *LocalEventBus) Unsubscribe(t string, fn domain.EventBusHandler) error {
	return e.bus.Unsubscribe(t, fn)
}

func NewLocalEventBus() domain.EventBus {
	// uncomment this if a singleton is needed
	// if localEventBus == nil {
	// 	localEventBus = &LocalEventBus{
	// 		bus: evbus.New(),
	// 	}
	// }

	return &LocalEventBus{
		bus: evbus.New(),
	}
}
