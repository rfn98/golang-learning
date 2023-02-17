package controller

import (
	"net/http"

	. "echo-go/src/model"
	. "echo-go/src/session"

	"github.com/labstack/echo"
)

type M map[string]interface{}

var store = NewCookieStore()
var SESSION_ID = "session_app_id"

func SayHello(c echo.Context) error {
	data := M{"Message": "Hello Veni Nur", "Counter": 10}
	return c.JSON(http.StatusOK, data)
}

func SetCookie(c echo.Context) error {
	session, _ := store.Get(c.Request(), SESSION_ID)
	session.Values["app_name"] = "HRIS"
	session.Values["app_key"] = "1234LKHHU)(&hj"
	session.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, "/get")
}

func GetUsers(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}
	if err = c.Validate(u); err != nil {
		return
	}
	/* for k, v := range *u { // GET VALUE OF TYPE DATA POINTER MAP INTERFACE
		fmt.Println(k, "=>", v)
	} */
	return c.JSON(http.StatusOK, u)
}

var ActionAbout = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("FROM ACTION ABOUT"))
		},
	),
)
