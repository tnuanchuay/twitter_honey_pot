package app

import (
	"net/http"
	"twitter_honey_pod/app/database"
)

func CreateApp() *http.Server{
	err := database.MigrateDb()
	if err != nil {
		panic(err)
	}

	return getServer()
}
