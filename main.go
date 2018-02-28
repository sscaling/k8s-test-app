package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func RootHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received %v\n", req)
	payload := fmt.Sprintf("OK - %s\n", time.Now().Format("Mon 01/02/06 03:04:05PM -07:00"))
	res.Write([]byte(payload))
}
