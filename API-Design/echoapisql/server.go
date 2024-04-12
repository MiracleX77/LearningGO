package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	user "mix/user"

	_ "github.com/lib/pq"
)

func main() {
	user.InitDB()
	e := echo.New()

	// e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	if username == "apidesign" || password == "45678" {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", user.CreateUserHandler)
	e.GET("/users", user.GetUserHandler)
	e.GET("/users/:id", user.GetUsersHandler)

	log.Fatal(e.Start(":2566"))
}
