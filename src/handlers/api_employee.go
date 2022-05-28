package handlers

import (
	"github.com/garfada/garfada/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateEmployee - Create a employee
func (c *Container) CreateEmployee(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DeleteEmployee - Delete a employee
func (c *Container) DeleteEmployee(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetImployees - Get employees
func (c *Container) GetImployees(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// UpdateEmployee - Update a employee
func (c *Container) UpdateEmployee(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
