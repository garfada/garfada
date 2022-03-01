package dish

type ModelRequest struct {
	Name        string
	Description string
}

type ModelResponse struct {
	ID string
	ModelRequest
}
