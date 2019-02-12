package main

import (
	"flag"
	"log"

	"github.com/ilgooz/service-http-server/httpserver"
	"github.com/mesg-foundation/core/client/service"
	"github.com/mesg-foundation/core/x/xsignal"
)

var (
	serverAddr = flag.String("serverAddr", ":2300", "Server's listening address")
)

func main() {
	flag.Parse()

	s, err := service.New()
	if err != nil {
		log.Fatal(err)
	}

	hs, err := httpserver.New(s, *serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	// start the http server service.
	go func() {
		log.Println("http server service has been started")
		log.Printf("http server listening at %s\n", hs.ListeningAddr)

		if err := hs.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	// wait for interrupt and gracefully shutdown the http server service.
	<-xsignal.WaitForInterrupt()

	log.Println("shutting down...")

	if err := hs.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("shutdown")
}
