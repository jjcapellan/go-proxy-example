package main

import (
	"log"
	"sync"

	backend "github.com/jjcapellan/go-proxy-example/backend"
	frontend "github.com/jjcapellan/go-proxy-example/frontend"
	proxy "github.com/jjcapellan/go-proxy-example/proxy"
)

func main() {
	var wg sync.WaitGroup

	go backend.Init()
	go frontend.Init()
	go proxy.Init()

	wg.Add(3)
	wg.Wait()

	log.Println("App terminated")
}
