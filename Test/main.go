package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/miracle/go-test/mix"
)

func main() {
	id := uuid.New()
	fmt.Println("Hello, World!s")
	fmt.Printf("%s", id)
	mix.SayHello()
}
