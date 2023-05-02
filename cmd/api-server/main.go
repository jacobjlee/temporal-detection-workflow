package main

import (
	"log"

	"github.com/jacobjlee/temporal-detection-workflow/internal/server"
	"go.temporal.io/sdk/client"
)

func main() {
	// create the temporal client
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	// create the server
	s := server.NewServer(temporalClient)
	s.MountHandlers()
	log.Fatal(s.Run("localhost", 8000))
}
