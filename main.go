package main

import "github.com/garfada/garfada/internal"

func main() {
	err := internal.Server("3000")
	if err != nil {
		return
	}
}
