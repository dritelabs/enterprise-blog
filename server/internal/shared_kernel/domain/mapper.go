package domain

type Mapper[E, D any] interface {
	ToDomain(e E) D
	ToEntity(d D) E
}
