package frontend

import (
	"log"
	"net/http"
)

func Init() {
	fs := http.FileServer(http.Dir("./frontend/public"))
	mux := http.NewServeMux()
	mux.Handle("/", fs)

	log.Println("Web server listening on port 3001")
	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		log.Fatal("Web server not started")
	}
}
