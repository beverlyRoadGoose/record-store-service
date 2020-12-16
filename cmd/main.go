package main

import (
	"flag"
	"heytobi.dev/record-store-service/internal"
	"log"
	"net/http"

	"heytobi.dev/record-store-service/internal/transport"
)

func main() {
	var (
		service     = internal.NewRecordStoreService()
		endpoints   = transport.MakeEndpoints(service)
		httpAddress = flag.String("http.address", ":8080", "HTTP address")
		httpHandler = transport.NewHTTPHandler(endpoints)
	)

	flag.Parse()
	server := &http.Server{
		Addr:    *httpAddress,
		Handler: httpHandler,
	}

	log.Printf("Starting up server on %s", *httpAddress)
	_ = server.ListenAndServe()
}
