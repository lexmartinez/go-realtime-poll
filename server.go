package main

import (
	"net/http"
)

func home (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func vote (w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	option := r.Form.Get("option")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Success vote for: " + option))
}

func base(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		vote(w,r)
	} else if r.Method == http.MethodGet {
		home(w,r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Unsupported Request Method!"))
	}
}

func results(w http.ResponseWriter, r *http.Request) {
	message := "results.html"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", base)
	http.HandleFunc("/results", results)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
