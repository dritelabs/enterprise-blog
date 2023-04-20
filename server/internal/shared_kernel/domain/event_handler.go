package domain

type EventHandler[E any] interface {
	Handle(e E) error
}
