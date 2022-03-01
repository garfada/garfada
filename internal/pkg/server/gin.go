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

	s.POST(business.Route, dishRoutes.Create)
	s.GET(business.Route, dishRoutes.Read)
	s.PATCH(business.Route, dishRoutes.Update)
	s.DELETE(business.Route, dishRoutes.Delete)

	s.POST(business.Route, employeeRoutes.Create)
	s.GET(business.Route, employeeRoutes.Read)
	s.PATCH(business.Route, employeeRoutes.Update)
	s.DELETE(business.Route, employeeRoutes.Delete)

	s.POST(business.Route, mealRoutes.Create)
	s.GET(business.Route, mealRoutes.Read)
	s.PATCH(business.Route, mealRoutes.Update)
	s.DELETE(business.Route, mealRoutes.Delete)

	err := s.Run(port)
	if err != nil {
		return err
	}

	return nil
}
