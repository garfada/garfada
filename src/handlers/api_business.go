package handlers

import (
	"github.com/garfada/garfada/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateBusiness - Create a business
func (c *Container) CreateBusiness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DeleteBusiness - Delete a business
func (c *Container) DeleteBusiness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetBusiness - Get businesses
func (c *Container) GetBusiness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// UpdateBusiness - Update a business
func (c *Container) UpdateBusiness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
