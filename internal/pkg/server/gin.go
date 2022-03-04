package server

import (
	"github.com/garfada/garfada/internal/business"
	"github.com/garfada/garfada/internal/dish"
	"github.com/garfada/garfada/internal/employee"
	"github.com/garfada/garfada/internal/meal"
	"github.com/gin-gonic/gin"
)

func NewGinServer(port string) error {
	businessRoutes := business.GinRoutes{}
	dishRoutes := dish.GinRoutes{}
	employeeRoutes := employee.GinRoutes{}
	mealRoutes := meal.GinRoutes{}

	s := gin.Default()

	s.POST(business.Route, businessRoutes.Create)
	s.GET(business.Route, businessRoutes.Read)
	s.PATCH(business.Route, businessRoutes.Update)
	s.DELETE(business.Route, businessRoutes.Delete)

	s.POST(dish.Route, dishRoutes.Create)
	s.GET(dish.Route, dishRoutes.Read)
	s.PATCH(dish.Route, dishRoutes.Update)
	s.DELETE(dish.Route, dishRoutes.Delete)

	s.POST(employee.Route, employeeRoutes.Create)
	s.GET(employee.Route, employeeRoutes.Read)
	s.PATCH(employee.Route, employeeRoutes.Update)
	s.DELETE(employee.Route, employeeRoutes.Delete)

	s.POST(meal.Route, mealRoutes.Create)
	s.GET(meal.Route, mealRoutes.Read)
	s.PATCH(meal.Route, mealRoutes.Update)
	s.DELETE(meal.Route, mealRoutes.Delete)

	err := s.Run("127.0.0.1:" + port)
	if err != nil {
		return err
	}

	return nil
}
