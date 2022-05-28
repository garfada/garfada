package handlers

import (
	"github.com/garfada/garfada/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateDish - Create a dish
func (c *Container) CreateDish(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DeleteDish - Delete a dish
func (c *Container) DeleteDish(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetDishes - Get dishes
func (c *Container) GetDishes(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// UpdateDish - Update a dish
func (c *Container) UpdateDish(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
