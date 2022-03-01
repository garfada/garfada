package meal

type ModelRequest struct {
	Name        string
	Description string
}

type ModelResponse struct {
	ID int
	ModelRequest
}
