package app

import (
	"net/http"
	"time"
)

func getServer() *http.Server {
	return &http.Server{
		Handler:      getRouter(),
		Addr:         ":8000",
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}
}
