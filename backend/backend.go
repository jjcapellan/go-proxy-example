package backend

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Ab struct {
	A int
	B int
}

var ab Ab = Ab{12, 24}
var apiKey string

// INIT FUNCTIONS

func Init(port string) {
	loadEnv()
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("Backend: error -> no api key found")
	}

	setupServer(port)
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Backend: environment file not loaded")
	}
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
	rt.Use(apiKeyMiddleware)
}

// Routes Handlers

func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xApiKey := r.Header.Get("x-api-key")
		if apiKey != xApiKey {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

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
