package handler

import (
	"log"
	"net/http"
	"twitter_honey_pod/app/service"
)

func CatchHandler (w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	referer := r.Header.Get("Referer")
	path := r.URL.Path

	_, err := service.StoreNewCatch(path, ip, referer, xForwardedFor)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "https://www.twitter.com", http.StatusTemporaryRedirect)
		return
	}

	h, err := service.FindHoneyByPath(path)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "https://www.twitter.com", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, h.RedirectTo, http.StatusTemporaryRedirect)
}
