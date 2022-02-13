package backend

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Ab struct {
	A int
	B int
}

var ab Ab = Ab{12, 24}

// INIT FUNCTIONS

func Init(port string) {
	setupServer(port)
}

func setupServer(port string) {
	rt := mux.NewRouter().StrictSlash(true)
	setupRoutes(rt)
	log.Printf("Backend server listening on port %s", port)
	err := http.ListenAndServe(":"+port, rt)
	if err != nil {
		log.Fatal("Api server not started")
	}
}

// ROUTES ////

func setupRoutes(rt *mux.Router) {
	rt.HandleFunc("/api/ab", abHandler).Methods("GET")
	rt.HandleFunc("/", homeHandler)
}

// Routes Handlers

func abHandler(w http.ResponseWriter, r *http.Request) {
	abJson, _ := json.Marshal(&ab)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(abJson)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Resource not found"))
}
