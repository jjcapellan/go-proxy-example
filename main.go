package main

import (
	"os"
	"strconv"
	"sync"

	backend "github.com/jjcapellan/go-proxy-example/backend"
	frontend "github.com/jjcapellan/go-proxy-example/frontend"
	proxy "github.com/jjcapellan/jjc-reverse-proxy"
)

const PORT_LIMIT int = 65535

var PORT_PROXY string // From environment
var PORT_FRONTEND string
var PORT_API string

const ROUTE_API string = "api"
const ROUTE_FRONTEND string = "/"

func main() {

	setupPorts()

	var wg sync.WaitGroup
	wg.Add(1)

	go backend.Init(PORT_API)
	go frontend.Init(PORT_FRONTEND)

	proxy.AddProxy(ROUTE_API, PORT_API)
	proxy.AddProxy(ROUTE_FRONTEND, PORT_FRONTEND)
	go func() {
		defer wg.Done()
		proxy.Start(PORT_PROXY)
	}()

	wg.Wait()
}

func setupPorts() {

	PORT_PROXY = getPortProxy()

	p, _ := strconv.Atoi(PORT_PROXY)
	if p > (PORT_LIMIT - 3) {
		p = p - 3
	}

	PORT_FRONTEND = strconv.Itoa(p + 1)
	PORT_API = strconv.Itoa(p + 2)
}

func getPortProxy() string {
	var port string
	if os.Getenv("PORT") == "" {
		port = "8080"
	} else {
		port = os.Getenv("PORT")
	}
	return port
}
