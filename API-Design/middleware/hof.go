package main

import "fmt"

type Decorator func(s string) error

func Use(next Decorator) Decorator {
	return func(c string) error {
		fmt.Println("Do something before the next")
		return next(c)
	}
}
func home(s string) error {
	fmt.Println("Home:", s)
	return nil
}

func main1() {
	warpped := Use(home)
	w := warpped("world")
	fmt.Println(w)
}
