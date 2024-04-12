package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Err struct {
	Message string `json:"message"`
}

var users = []User{
	{ID: 1, Name: "AnuchitO", Age: 18},
}

func createUserHandler(c echo.Context) error {
	u := User{}
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	users = append(users, u)

	fmt.Println("id : % #v\n", u)

	return c.JSON(http.StatusCreated, u)
}

func getUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func AuthMiddleware(username, password string, c echo.Context) (bool, error) {
	if username == "apidesign" || password == "45678" {
		return true, nil
	}
	return false, nil
}

type jwtCustomClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username != "join" || password != "123" {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		"Jon Snow",
		"admin",
		"accessToken",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 60)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	refrechTokenClaims := &jwtCustomClaims{
		"Jon Snow",
		"admin",
		"refreshToken",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refrechTokenClaims,
	})
}

func refreshToken(c echo.Context) error {
	refreshTokenString := c.FormValue("refresh_token")
	jwtSecret := []byte("secret")
	token, err := jwt.ParseWithClaims(refreshTokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return echo.ErrUnauthorized
	}
	claims, ok := token.Claims.(*jwtCustomClaims)
	if !ok || !token.Valid {
		return echo.ErrUnauthorized
	}
	if claims.Type != "refreshToken" {
		return echo.ErrUnauthorized
	}
	accessTokenClaims := &jwtCustomClaims{
		claims.Name,
		claims.Role,
		"accessToken",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 60)),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  accessTokenString,
		"refreshToken": refreshTokenString,
	})
}

func jwtMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return echo.ErrUnauthorized
		}
		parts := strings.Split(authToken, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			return echo.ErrUnauthorized
		}
		jwtToken := parts[1]
		token, err := jwt.ParseWithClaims(jwtToken, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*jwtCustomClaims)
		if !ok || !token.Valid {
			return echo.ErrUnauthorized
		}
		c.Set("user", claims.Name)
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	g := e.Group("/api")
	g.POST("/login", login)
	g.POST("/refresh", refreshToken)

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	g.Use(echojwt.WithConfig(config))
	//g.Use(middleware.BasicAuth(AuthMiddleware))
	g.Use(jwtMiddleWare)
	g.POST("/users", createUserHandler)
	g.GET("/users", getUsersHandler)

	log.Fatal(e.Start(":2565"))
}
