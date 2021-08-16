package frontend

import (
	"log"
	"net/http"
)

func Init(port string) {
	fs := http.FileServer(http.Dir("./frontend/public"))
	mux := http.NewServeMux()
	mux.Handle("/", fs)

	log.Printf("Web server listening on port %s", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal("Web server not started")
	}
}
