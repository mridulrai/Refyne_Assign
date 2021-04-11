package rest

import (
	"github.com/labstack/echo/v4"
)

// InitRoutes initializes routes
func InitRoutes() *echo.Echo {
	e := echo.New()

	e.GET("/", index)
	e.GET("/health", health)
	e.POST("/v1/user", createUser)
	e.POST("/v1/cars", addCars)
	e.GET("/v1/searchCars", listAccounts)            //contains query from given timeDate to given timeDate, returns the list of avialable cars
	e.GET("/v1/calculatePrice", calculatePrice)      //contains query  from given timeDate to given timeDate
	e.GET("/v1/user/:id/bookings", listUserBookings) //paticular user booking details
	e.GET("/v1/cars/:id/bookings", listCarBookings)  //paticular car booking details
	e.POST("/v1/cars/:id/book", bookCar)

	return e
}
