package business

type ModelRequest struct {
	Code string
	Name string
}

type ModelResponse struct {
	ID int
	ModelRequest
}
