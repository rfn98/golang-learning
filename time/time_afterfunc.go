package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		ch <- true
		fmt.Println("EXPIRED")
	})
	fmt.Println("start")
	<- ch // WAITING 4 second
	fmt.Println("finish")
}