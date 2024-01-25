package main

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Data string `json:"data"`
}

type Pubsub struct {
	subs []chan Message
	mu   sync.Mutex
}

func (ps *Pubsub) Subscribe() chan Message {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(chan Message)
	ps.subs = append(ps.subs, ch)
	return ch
}

func (ps *Pubsub) Publish(msg *Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for _, ch := range ps.subs {
		ch <- *msg
	}
}
func (ps *Pubsub) Close(ch chan Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for i, sub := range ps.subs {
		if sub == ch {
			ps.subs = append(ps.subs[:i], ps.subs[i+1:]...)
			close(ch)
			return
		}
	}
}
func main() {
	app := fiber.New()

	pubsub := &Pubsub{}
	app.Post("/publisher", func(c *fiber.Ctx) error {
		message := new(Message)
		if err := c.BodyParser(message); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		pubsub.Publish(message)
		return c.JSON(&fiber.Map{
			"message": "success",
		})
	})

	sub := pubsub.Subscribe()
	go func() {
		for msg := range sub {
			fmt.Println(msg)
		}
	}()
	app.Listen(":3000")
}
