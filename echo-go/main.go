package main

import (
	. "echo-go/src/helper"
	. "echo-go/src/log"
	. "echo-go/src/middleware"
	. "echo-go/src/router"
	. "echo-go/src/session"
	"fmt"
	"time"

	"github.com/labstack/echo"
)

func main() {
	r := echo.New()
	r.Use(MiddlewareLogging)
	/* r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	})) */
	fmt.Println("WKWKWKWWK", NewCookieStore())
	Validate(r)
	IndexRouter(r)

	lock := make(chan error)
	go func(lock chan error) {
		lock <- r.Start(":4321")
	}(lock)

	time.Sleep(1 * time.Millisecond)
	MakeLogEntry(nil).Warning("application started without ssl/tls enabled")

	err := <-lock
	if err != nil {
		MakeLogEntry(nil).Panic("failed to start application")
	}
	// r.Start(":4321")
}
