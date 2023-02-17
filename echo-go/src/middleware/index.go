package middleware

import (
	. "echo-go/src/log"
	"fmt"

	"github.com/labstack/echo"
)

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		MakeLogEntry(ctx).Info("INCOMING REQUEST")
		return next(ctx)
	}
}

func MiddlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("FROM MIDDLEWARE ONE")
		return next(ctx)
	}
}

func MiddlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("FROM MIDDLEWARE TWO")
		return next(ctx)
	}
}
