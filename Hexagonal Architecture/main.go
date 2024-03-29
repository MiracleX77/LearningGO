package main

import (
	"fmt"
	"log"
	"mix/adapters"
	"mix/core"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	database = "mydatabase"
	username = "myuser"
	password = "mypassword"
)

// 	"mix/core"

// 	"github.com/gofiber/fiber/v2"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"

func main() {
	app := fiber.New()
	// fmt.Println("Hello, playground")
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
	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)

	db.AutoMigrate(&core.Order{})

	app.Listen(":3000")
}
