package main

import (
	"fmt"
	"time"
)

func main() {
	<-time.After(4 * time.Second)
	fmt.Println("EXPIRED")	
}