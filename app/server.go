package app

import (
	"net/http"
	"os"
	"time"
)

func getServer() *http.Server {
	return &http.Server{
		Handler:      getRouter(),
		Addr:         os.Getenv("PORT"),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
}
