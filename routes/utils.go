package routes

import (
	"net/http"
	"strconv"
)

func CORS(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		next(w, r)
	})
}

func GetPropsFromURL(r *http.Request) (int, int, int) {

	var page int
	var idtype int
	var idability int
	var err error

	if len(r.URL.Query()["page"]) > 0 {
		page, err = strconv.Atoi(r.URL.Query()["page"][0])
		if err != nil {
			page = 1
		}
	} else {
		page = 1
	}

	if len(r.URL.Query()["type"]) > 0 {
		idtype, err = strconv.Atoi(r.URL.Query()["type"][0])
		if err != nil {
			idtype = 0
		}
	} else {
		idtype = 0
	}

	if len(r.URL.Query()["ability"]) > 0 {
		idability, err = strconv.Atoi(r.URL.Query()["ability"][0])
		if err != nil {
			idability = 0
		}
	} else {
		idability = 0
	}

	return page, idtype, idability

}
