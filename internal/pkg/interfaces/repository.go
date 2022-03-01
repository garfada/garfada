package interfaces

type Repository[Q any, P any] interface {
	Save(model Q) (*P, error)
	Restore(id int) (*P, error)
	Update(id int, model Q) (*P, error)
	Remove(id int) (*P, error)
}
