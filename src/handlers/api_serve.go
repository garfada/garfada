package handlers

import (
	"github.com/garfada/garfada/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ServeMeal - Assign a meal to a employee
func (c *Container) ServeMeal(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
