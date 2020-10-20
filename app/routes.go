package app

import (
	"github.com/gorilla/mux"
	"twitter_honey_pod/app/handler"
)

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/honey", handler.CreateHoneyHandler).Methods("POST")
	r.HandleFunc("/honey", handler.GetHoneyHandler).Methods("GET")
	return r
}
