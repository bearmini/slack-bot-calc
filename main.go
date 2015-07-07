package main

import (
	"net/http"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", top)
	m.Post("/", top)
	m.Run()
}

func top(w http.ResponseWriter, r *http.Request) (int, string) {
	if err := r.ParseForm(); err != nil {
		return 500, "Need parameters";
	}

	result := ""
	for k, v := range r.Form {
		for _, s := range v {
			result += k + ": " + s + "\n"
		}
	}

	return 200, result
}
