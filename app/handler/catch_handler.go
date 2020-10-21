package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"twitter_honey_pod/app/model"
	"twitter_honey_pod/app/service"
)

func CreateCatchHandler(w http.ResponseWriter, r *http.Request) {
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

func GetCatchHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	path := v["path"]

	catches, err := service.GetCatchByPath(path)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, err.Error())
		return
	}

	format := r.URL.Query().Get("format")
	var out string
	switch format {
	case "json":
		b, err := json.Marshal(catches)
		if err != nil {
			log.Println(err)
			fmt.Fprint(w, err.Error())
			return
		}

		out = string(b)
	default:
		out = BuildHtml(catches)
	}

	fmt.Fprint(w, out)
}

func BuildHtml(catches []model.Catch) string {
	var out string
	for _, elem := range catches {
		out = fmt.Sprintf("%s<br>%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%f\t%f\t%s\t%s\t%s\t%s\t%s\t%t\t%t", out,
			elem.Id, elem.HitTime, elem.Ip, elem.RefererUrl, elem.XForwardedFor,
			elem.City, elem.CountryName, elem.CountryCode, elem.ContinentName,
			elem.Latitude, elem.Longitude, elem.AsnId, elem.AsnName, elem.AsnDomain,
			elem.AsnRoute, elem.AsnType, elem.IsTor, elem.IsProxy)
		out = fmt.Sprintf("%s<a href=\"https://www.google.co.th/maps/@%f,%f,16.75z?hl=th\">link</a>", out, elem.Latitude, elem.Longitude)
	}

	return out
}
