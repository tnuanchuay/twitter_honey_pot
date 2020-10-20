package main

import (
	"log"
	"twitter_honey_pod/app"
)

func main () {
	log.Fatal(app.CreateApp().ListenAndServe())
}
