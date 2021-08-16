package main

import (
	"os"
	"strconv"
	"sync"

	backend "github.com/jjcapellan/go-proxy-example/backend"
	frontend "github.com/jjcapellan/go-proxy-example/frontend"
	proxy "github.com/jjcapellan/go-proxy-example/proxy"
)

const PORT_LIMIT int = 65535

var PORT_PROXY string
var PORT_FRONTEND string
var PORT_BACKEND string

func main() {

	setupPorts()

	var wg sync.WaitGroup

	go backend.Init(PORT_BACKEND)
	go frontend.Init(PORT_FRONTEND)
	go proxy.Init(PORT_PROXY, PORT_FRONTEND, PORT_BACKEND)

	wg.Add(1)
	wg.Wait()
}

func setupPorts() {

	PORT_PROXY = getPortProxy()

	p, _ := strconv.Atoi(PORT_PROXY)
	if p > (PORT_LIMIT - 3) {
		p = p - 3
	}

	PORT_FRONTEND = strconv.Itoa(p + 1)
	PORT_BACKEND = strconv.Itoa(p + 2)
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
