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

func setupRoutes(rt *mux.Router) {
	rt.HandleFunc("/api/ab", abHandler).Methods("GET")
	rt.HandleFunc("/", homeHandler)
}

func Init() {
	rt := mux.NewRouter().StrictSlash(true)
	setupRoutes(rt)
	log.Println("Serving backend on port 3002...")
	err := http.ListenAndServe(":3002", rt)
	if err != nil {
		log.Fatal(err)
	}
}
