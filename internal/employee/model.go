package employee

type ModelRequest struct {
	Code     string
	Name     string
	Business int
}

type ModelResponse struct {
	ID int
	ModelRequest
}
