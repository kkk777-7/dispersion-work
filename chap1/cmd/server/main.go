package main

import (
	"log"

	"github.com/kkk777-7/dispersion/chap1/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
