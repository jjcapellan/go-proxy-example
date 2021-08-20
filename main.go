package main

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"

	backend "github.com/jjcapellan/go-proxy-example/backend"
	frontend "github.com/jjcapellan/go-proxy-example/frontend"
	proxy "github.com/jjcapellan/jjc-reverse-proxy"
)

const PORT_LIMIT int = 65535

var PORT_PROXY string // From environment
var PORT_FRONTEND string
var PORT_API string

const API_ROUTE string = "/api"

func main() {

	loadEnvironment()
	setupPorts()

	var wg sync.WaitGroup
	wg.Add(1)

	go backend.Init(PORT_API)
	go frontend.Init(PORT_FRONTEND)

	proxyConfig := getProxyConfig()
	proxy.Setup(proxyConfig)
	go func() {
		defer wg.Done()
		proxy.Start()
	}()

	wg.Wait()
}

func loadEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Main: environment file not found")
	}
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

func getProxyConfig() proxy.Config {
	c := proxy.Config{
		PortProxy:    PORT_PROXY,
		PortApi:      PORT_API,
		PortFrontend: PORT_FRONTEND,
		ApiRoute:     API_ROUTE,
		ApiKey:       os.Getenv("API_KEY"),
	}
	return c
}
