package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	database = "mydatabase"
	username = "myuser"
	password = "mypassword"
)

func authRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	secretKey := "secret"
	if cookie == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	if !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claim := token.Claims.(jwt.MapClaims)
	fmt.Println(claim["user_id"])
	return c.Next()
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}
	print(db)

	fmt.Println("Connection Opened to Database")

	db.AutoMigrate(&Book{}, &User{})

	app := fiber.New()

	app.Use("/books", authRequired)

	app.Get("/books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		books := getBook(db, uint(id))
		return c.JSON(books)
	})
	app.Post("/register", func(c *fiber.Ctx) error {
		if !c.Is("json") {
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}
		if c.Body() == nil {
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if user.Email == "" || user.Password == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Email or Password is missing")
		}

		err = createUser(db, user)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.SendStatus(fiber.StatusCreated)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		if !c.Is("json") {
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}
		if c.Body() == nil {
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if user.Email == "" || user.Password == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Email or Password is missing")
		}

		token, err := loginUser(db, user)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 72),
			HTTPOnly: true,
		})
		return c.JSON(fiber.Map{
			"Message": "Login Successful",
		})
	})

	app.Listen(":8080")
}
