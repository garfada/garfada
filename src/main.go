package main

import (
	"github.com/garfada/garfada/src/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//TODO: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CreateBusiness - Create a business
	e.POST("/v1/businesses", c.CreateBusiness)

	// DeleteBusiness - Delete a business
	e.DELETE("/v1/businesses", c.DeleteBusiness)

	// GetBusiness - Get businesses
	e.GET("/v1/businesses", c.GetBusiness)

	// UpdateBusiness - Update a business
	e.PUT("/v1/businesses", c.UpdateBusiness)

	// CreateDish - Create a dish
	e.POST("/v1/dishes", c.CreateDish)

	// DeleteDish - Delete a dish
	e.DELETE("/v1/dishes", c.DeleteDish)

	// GetDishes - Get dishes
	e.GET("/v1/dishes", c.GetDishes)

	// UpdateDish - Update a dish
	e.PUT("/v1/dishes", c.UpdateDish)

	// CreateEmployee - Create a employee
	e.POST("/v1/employees", c.CreateEmployee)

	// DeleteEmployee - Delete a employee
	e.DELETE("/v1/employees", c.DeleteEmployee)

	// GetImployees - Get employees
	e.GET("/v1/employees", c.GetImployees)

	// UpdateEmployee - Update a employee
	e.PUT("/v1/employees", c.UpdateEmployee)

	// CreateMeal - Create a meal
	e.POST("/v1/meals", c.CreateMeal)

	// DeleteMeal - Delete a meal
	e.DELETE("/v1/meals", c.DeleteMeal)

	// GetMeals - Get meals
	e.GET("/v1/meals", c.GetMeals)

	// UpdateMeal - Update a meal
	e.PUT("/v1/meals", c.UpdateMeal)

	// ServeMeal - Assign a meal to a employee
	e.POST("/v1/serve", c.ServeMeal)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
