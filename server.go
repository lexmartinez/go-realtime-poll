package main

import (
	"net/http"
	"github.com/Joker/jade"
	"io/ioutil"
	"fmt"
	"github.com/pusher/pusher-http-go"
	"github.com/alexsasharegan/dotenv"
	"os"
	"html/template"
)


type Keys struct {
	Key   string
	Cluster string
}

func home (w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("./templates/index.go")
	if err != nil {
		fmt.Printf("ReadFile error: %v", err)
		return
	}

	tmpl, err := jade.Parse("index", string(dat))
	if err != nil {
		fmt.Printf("Parse error: %v", err)
		return
	}

	w.Write([]byte(tmpl))
}

func vote (w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	option := r.Form.Get("option")

	client := pusher.Client{
		AppId: os.Getenv("APP_ID"),
		Key: os.Getenv("APP_KEY"),
		Secret: os.Getenv("APP_SECRET"),
		Cluster: os.Getenv("APP_CLUSTER"),
		Secure: true,
	}

	data := map[string]string{"vote": option}
	client.Trigger("tv-shows", "vote-event", data)

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
	dat, err := ioutil.ReadFile("./templates/results.go")
	if err != nil {
		fmt.Printf("ReadFile error: %v", err)
		return
	}

	tmpl, err := jade.Parse("results", string(dat))
	if err != nil {
		fmt.Printf("Parse error: %v", err)
		return
	}

	keys := Keys{
		Key:   os.Getenv("APP_KEY"),
		Cluster: os.Getenv("APP_CLUSTER"),
	}

	goTpl, err := template.New("html").Parse(tmpl)
	if err != nil {
		fmt.Printf("\nTemplate parse error: %v", err)
		return
	}
	err = goTpl.Execute(w, keys)
	if err != nil {
		fmt.Printf("\nExecute error: %v", err)
		return
	}
}

func chartjs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/chart.js")
}

func main() {

	err := dotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}
	http.HandleFunc("/", base)
	http.HandleFunc("/results", results)
	http.HandleFunc("/chart.js", chartjs)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	} else {
		fmt.Println("Polling server running at localhost:8080")
	}
}
