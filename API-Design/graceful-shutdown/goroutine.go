package main

import (
	"fmt"
	"time"
)

func slow(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Second)
	}
}
func main0() {

	done := make(chan bool)
	go func() {
		slow("Hello")
		done <- true
	}()
	time.Sleep(3 * time.Second)
	<-done
	fmt.Println("Done")
}
