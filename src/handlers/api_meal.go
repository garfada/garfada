package handlers

import (
	"github.com/garfada/garfada/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateMeal - Create a meal
func (c *Container) CreateMeal(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DeleteMeal - Delete a meal
func (c *Container) DeleteMeal(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetMeals - Get meals
func (c *Container) GetMeals(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// UpdateMeal - Update a meal
func (c *Container) UpdateMeal(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
