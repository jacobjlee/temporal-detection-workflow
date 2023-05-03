package main

import (
	"log"

	"go.temporal.io/sdk/client"

	"github.com/jacobjlee/temporal-detection-workflow/internal/detection"
	"github.com/jacobjlee/temporal-detection-workflow/internal/server"
)

func main() {
	// create the temporal client
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	detectionService := detection.NewDetectionService(detection.NewDetectionRepository())

	// create the server
	s := server.NewServer(temporalClient, detectionService)
	s.MountHandlers()
	log.Fatal(s.Run("localhost", 8000))
}
