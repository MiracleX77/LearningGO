package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	database = "mydatabase"
	username = "myuser"
	password = "mypassword"
)

var db *sql.DB

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)

	sdb, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	db = sdb
	defer db.Close()
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/product/:id", getProductHandler)
	app.Post("/product", createProductHandler)
	app.Put("/product/:id", updateProductHandler)
	app.Get("/products", getProductsHandler)
	app.Listen(":8080")

}

func getProductHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid product id")
	}
	product, err := getProduct(id)
	if err != nil {
		return c.Status(404).SendString("No product found with id " + c.Params("id"))
	}
	return c.JSON(product)
}
func createProductHandler(c *fiber.Ctx) error {
	p := new(Product)
	err := c.BodyParser(p)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	createProduct(p)
	return c.JSON(p)
}
func updateProductHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid product id")
	}
	p := new(Product)
	err = c.BodyParser(p)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = updateProduct(id, p)
	if err != nil {
		return c.Status(404).SendString("No product found with id " + c.Params("id"))
	}
	return c.JSON(p)
}

func getProductsHandler(c *fiber.Ctx) error {
	products, err := getProducts()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(products)
}
