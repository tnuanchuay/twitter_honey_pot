package app

import (
	"github.com/gorilla/mux"
	"twitter_honey_pod/app/handler"
)

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/catch/{path}", handler.GetCatchHandler).Methods("GET")
	r.HandleFunc("/honey", handler.CreateHoneyHandler).Methods("POST")
	r.HandleFunc("/{path}", handler.CreateCatchHandler).Methods("GET")
	return r
}
