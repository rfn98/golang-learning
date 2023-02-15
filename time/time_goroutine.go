package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
)

func timer(timeout int, ch chan<- bool /*SEND*/) {
	time.AfterFunc(time.Duration(timeout) * time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool /*RECEIVE*/) {
	time.AfterFunc(time.Duration(timeout) * time.Second, func() {
		<-ch
	    fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	    os.Exit(1)
	})
}

func main () {
	var (
		timeout = 5
		ch = make(chan bool)
	)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Print("What is 725/25 ? ")
	fmt.Scan(&input)

	if input == strconv.Itoa(725/25) {
		fmt.Println("Right Answer!")
	} else {
		fmt.Println("Wrong Answer!")
	}
}