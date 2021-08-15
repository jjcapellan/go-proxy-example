package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var backend *httputil.ReverseProxy
var frontend *httputil.ReverseProxy

func routesHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.String(), "/api/") {
		backend.ServeHTTP(w, r)
		return
	}
	frontend.ServeHTTP(w, r)
}

func Init() {

	var err error
	backend, err = newProxy("http://localhost:3002")
	if err != nil {
		log.Fatal("Proxy: Error creating backend handler")
	}
	frontend, err = newProxy("http://localhost:3001")
	if err != nil {
		log.Fatal("Proxy: Error creating frontend handler")
	}

	http.HandleFunc("/", routesHandler)

	log.Println("Proxy running on port 3000")
	err = http.ListenAndServe(":3000", nil)
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
