package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"twitter_honey_pod/app/service"
)

func CreateHoneyHandler (w http.ResponseWriter, r *http.Request) {
	var userId = r.FormValue("userId")
	var redirectTo = r.FormValue("redirectTo")
	var url = r.FormValue("url")

	result, err := service.StoreNewHoney(userId, redirectTo, url)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(*result)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, string(b))
}