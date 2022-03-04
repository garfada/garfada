package internal

import "github.com/garfada/garfada/internal/pkg/server"

func Server(port string) error {
	err := server.NewGinServer(port)
	if err != nil {
		return err
	}

	return nil
}
