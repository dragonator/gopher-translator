package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dragonator/gopher-translator/internal/service"
)

func main() {
	_ = os.Setenv("SPECIFICATION_FILEPATH", "/Users/tdraganov/code/gopher-translator/configs/gopher_rules.json")

	config := map[string]string{}
	specFile := os.Getenv("SPECIFICATION_FILEPATH")
	if specFile == "" {
		log.Fatal("missing environment variable $SPECIFICATION_FILEPATH")
	}
	config[service.CfgParamSpecFile] = specFile
	config[service.CfgParamPortNumber] = "8080" // TODO: Retrieve from parameter

	svc, err := service.New(config)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	svc.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop

	log.Printf("Signal caught (%s), stopping...", sig.String())
	svc.Stop()
	log.Print("Service stopped.")
}
