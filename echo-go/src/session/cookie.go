package session

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

func NewCookieStore() *sessions.CookieStore {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	authKey := []byte(viper.GetString("appName"))
	encryptionKey := []byte(viper.GetString("secretKey"))

	store := sessions.NewCookieStore(authKey, encryptionKey)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7
	store.Options.HttpOnly = true

	return store
}
