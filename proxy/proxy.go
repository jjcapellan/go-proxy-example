package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var backend *httputil.ReverseProxy
var frontend *httputil.ReverseProxy
var apiKey string

func Init(port string, portFrontend string, portBackend string) {
	loadEnv()
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		apiKey = "noKey"
	}

	createProxies(portFrontend, portBackend)
	setupServer(port)
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Backend: environment file not loaded")
	}
}

func createProxies(portFrontend string, portBackend string) {
	var err error
	backend, err = newProxy("http://localhost:" + portBackend)
	if err != nil {
		log.Fatal("Proxy: Error creating backend handler")
	}
	frontend, err = newProxy("http://localhost:" + portFrontend)
	if err != nil {
		log.Fatal("Proxy: Error creating frontend handler")
	}
}

func setupServer(port string) {
	http.HandleFunc("/", routesHandler)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Proxy server listening on port ", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Proxy: Error initializing server")
	}
}

func newProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(url), nil
}

func routesHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.String(), "/api/") {
		r.Header.Set("x-api-key", apiKey)
		backend.ServeHTTP(w, r)
		return
	}
	frontend.ServeHTTP(w, r)
}
