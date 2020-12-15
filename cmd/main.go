package main

import (
	"flag"
	"log"
	"net/http"

	"heytobi.dev/record-store-service/pkg/recordstore"
	"heytobi.dev/record-store-service/pkg/recordstore/transport"
)

func main() {
	var (
		service     = recordstore.NewRecordStoreService()
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
