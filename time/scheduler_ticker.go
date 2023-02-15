package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(4 * time.Second) // WAITING 4 SECOND
		done <- true // SEND TRUE TO CHAN DONE
	}()

	for {
		select {
		case <-done: // CHAN DONE RECEIVE TRUE
			ticker.Stop()
			return
		case t := <-ticker.C: // GET DATE TIME EVENT HAPPENED
			fmt.Println("HELLO", t)
		}
	}
}