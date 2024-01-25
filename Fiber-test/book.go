package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Handler functions
// getBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} Book
// @Router /books [get]
func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookId := c.Params("id")
	id, err := strconv.Atoi(bookId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.Id == id {
			return c.JSON(book)
		}
	}
	return c.Status(404).SendString("No book found with IDs")
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	upBook := new(Book)
	if err := c.BodyParser(upBook); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	for i, book := range books {
		if book.Id == bookId {
			books[i].Title = upBook.Title
			books[i].Author = upBook.Author
			return c.JSON(books[i])
		}
	}
	return c.Status(404).SendString("No book found with IDs")

}
func deleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for i, book := range books {
		if book.Id == bookId {
			books = append(books[:i], books[i+1:]...)
			return c.SendString("Book is deleted")
		}
	}
	return c.Status(404).SendString("No book found with IDs")
}
