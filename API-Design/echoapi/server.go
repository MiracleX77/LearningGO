package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "AnuchitO", Age: 18},
}

// echo

type Err struct {
	Message string `json:"message"`
}

func getUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}
func createUserHandler(c echo.Context) error {
	u := User{}
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	users = append(users, u)
	return c.JSON(http.StatusCreated, u)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/api")
	g.GET("/users", getUsersHandler)
	g.POST("/users", createUserHandler)

	log.Println("Server started at :2565")
	log.Fatal(e.Start(":2565"))
	log.Println("bye bye!")
}
