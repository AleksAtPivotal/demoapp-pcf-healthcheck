package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	appVersion = "1.0"
)

type App struct {
	crashTrigger int
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {

	var app App
	// crashTrigger is used to crash the app on X number of probes to healthz endpoint
	app.crashTrigger = 20

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		handleHealth(w, r, &app)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World version %s", appVersion)
	log.Printf("Hello World version %s", appVersion)
}

func handleHealth(w http.ResponseWriter, r *http.Request, a *App) {
	log.Printf("Recieved a probe to /healthz endpoint. 500 countdown : %v", a.crashTrigger)
	if a.crashTrigger == 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	} else {
		w.WriteHeader(http.StatusOK)
		a.crashTrigger = a.crashTrigger - 1
	}
}
