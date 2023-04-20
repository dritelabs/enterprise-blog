package domain

type QueryHandler[Q any] interface {
	Execute(q Q) error
}
