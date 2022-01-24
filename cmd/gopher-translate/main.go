package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/dragonator/gopher-translator/internal/service"
)

func main() {
	port := flag.Int("port", 8080, "port number the server should listen on")
	specfile := flag.String("specfile", "", "json file with translation instructions")
	flag.Parse()

	if *specfile == "" {
		log.Fatal("missing required argument: --specfile FILE")
	}
	f, err := os.Open(*specfile)
	if err != nil {
		log.Fatalf("failure opening spec file: %v", err)
	}

	b := &service.Bootstrap{
		Port: strconv.Itoa(*port),
		Spec: f,
	}
	svc, err := service.New(b)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
	// Close now to avoid keeping the file descriptor
	// open during the whole time running the service
	f.Close()

	svc.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop

	log.Printf("Signal caught (%s), stopping...", sig.String())
	svc.Stop()
	log.Print("Service stopped.")
}
