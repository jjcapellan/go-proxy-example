package main

import (
	"log"
	"net/http"
	"strings"

	backend "github.com/jjcapellan/go-proxy-example/backend"
	proxy "github.com/jjcapellan/go-proxy-example/proxy"
)

// backend proxy implements Handler interface
var backendProxy *proxy.Proxy

var fsHandler http.Handler = http.FileServer(http.Dir("./frontend/public"))

func routesHandler(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.URL.String(), "/api/") {
		backendProxy.ServeHTTP(w, r)
		return
	}

	fsHandler.ServeHTTP(w, r)
}

func main() {

	err := proxy.Init()
	if err != nil {
		log.Fatal("Backend proxy not created")
	}

	backendProxy = &proxy.BackendProxy

	go backend.Init()

	log.Println("Frontend listening on port 3000...")
	http.HandleFunc("/", routesHandler)
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Main server not started")
	}
}
