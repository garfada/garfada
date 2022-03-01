package business

type Repository struct{}

func (r *Repository) Save(model ModelRequest) (*ModelResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Restore(id int) (*ModelResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Update(id int, model ModelRequest) (*ModelResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Remove(id int) (*ModelResponse, error) {
	//TODO implement me
	panic("implement me")
}
