package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()

	time.Sleep(5 * time.Second)
	
	t2 := time.Now()

	duration := t2.Sub(t1)

	fmt.Println("Duration in seconds:", duration.Seconds())
	fmt.Println("Duration in minutes:", duration.Minutes())
	fmt.Println("Duration in hours:", duration.Hours())

	/*start := time.Now()

	time.Sleep(5 * time.Second)

	duration := time.Since(start)

	fmt.Println("Duration in seconds:", duration.Seconds())
	fmt.Println("Duration in minutes:", duration.Minutes())
	fmt.Println("Duration in hours:", duration.Hours())*/
}