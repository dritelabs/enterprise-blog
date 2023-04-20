package domain

type EventBusHandler func(e Event)

type EventBus interface {
	Publish(t string, e Event)
	Subscribe(t string, fn EventBusHandler) error
	Unsubscribe(t string, fn EventBusHandler) error
}
