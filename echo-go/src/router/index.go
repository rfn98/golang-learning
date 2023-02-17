package router

import (
	. "echo-go/src/controller"

	"github.com/labstack/echo"
)

func IndexRouter(r *echo.Echo) {
	r.GET("/json", SayHello)
	r.GET("/about", ActionAbout)
	r.POST("/users", GetUsers)
}
