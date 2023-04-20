package domain

type Repository[E any] interface {
	Get(id string) (*E, error)
	Save(e E) (*E, error)
	Update(e E) (*E, error)
}
