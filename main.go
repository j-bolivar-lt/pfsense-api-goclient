package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/j-bolivar-lt/pfsense-api-goclient/v2/pfsenseapi"
)

func main() {
	// Create a new client with local authentication
	client := pfsenseapi.NewClientWithLocalAuth("http://10.5.3.69:10443", "admin", "pfsense")

	// Set a timeout for the context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Try to get system version information
	version, err := client.System.GetVersion(ctx)
	if err != nil {
		log.Fatalf("Error getting system version: %v", err)
	}

	fmt.Printf("pfSense Version: %s\n", version.Version)
	fmt.Printf("Platform: %s\n", version.Platform)
	fmt.Printf("Architecture: %s\n", version.Architecture)

	// List interfaces
	interfaces, err := client.Interface.ListInterfaces(ctx)
	if err != nil {
		log.Fatalf("Error listing interfaces: %v", err)
	}

	fmt.Println("\nInterfaces:")
	for _, iface := range interfaces {
		fmt.Printf("- %s (%s)\n", iface.Id, iface.Descr)
	}

	fmt.Println("\nConnection to pfSense API successful!")
}
