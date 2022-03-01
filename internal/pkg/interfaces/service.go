package interfaces

type Service[Q any, P any] interface {
	Create(model Q) (*P, error)
	Read(id int) (*P, error)
	Update(id int, model Q) (*P, error)
	Delete(id int) (*P, error)
}
