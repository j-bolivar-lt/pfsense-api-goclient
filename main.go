package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/j-bolivar-lt/pfsense-api-goclient/v2/pfsenseapi"
	"os"
)

type DNSResolverHostOverrideAlias = pfsenseapi.DNSResolverHostOverrideAlias
type DNSResolverHostOverride = pfsenseapi.DNSResolverHostOverride

func main() {
	// Create a new client with local authentication
	password := os.Getenv("PFSENSE_PASSWORD")
	admin := os.Getenv("PFSENSE_ADMIN")
	pfhost := os.Getenv("PFSENSE_HOST")
	client := pfsenseapi.NewClientWithLocalAuth(pfhost, admin, password)

	// Set a timeout for the context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Try to get system version information
	version, err := client.System.GetVersion(ctx)
	if err != nil {
		log.Fatalf("Error getting system version: %v", err)
	}

	fmt.Printf("pfSense Version: %s\n", version.Version)

	system, err := client.System.GetSystemStatus(ctx)

	fmt.Println("System CPU Model:", system.Data.Cpu_model)
	fmt.Println("Uptime:", system.Data.Uptime)

	dns, err := client.System.GetDNSResolverSettings(ctx)
	if err != nil {
		log.Fatalf("Error getting DNS resolver settings: %v", err)
	}

	for _, server := range dns.Data.DnsServer {
		fmt.Printf("DNS Server: %s\n", server)
	}

	// dnsResponse, error := client.System.PatchDNSResolverSettings(ctx, dns.Data)

	// if error != nil {
	// 	log.Fatalf("Error patching DNS resolver settings: %v", error)
	// }

	// fmt.Println("DNS Resolver settings patched successfully!", dnsResponse)

	dnsHostResponse, error := client.DNS.GetDNSResolverHostOverrides(ctx)

	if error != nil {
		log.Fatalf("Error getting DNS resolver host overrides: %v", error)
	}

	for _, host := range dnsHostResponse {
		fmt.Printf("ID: %d\n", host.Id)
		fmt.Printf("Domain: %s\n", host.Domain)
		for key, ip := range host.IP {
			fmt.Printf("IP %d: %s\n", key, ip)
		}
	}

	DNSHostConfigArray := make([]DNSResolverHostOverrideAlias, 0)

	DNSHostConfigArray = append(DNSHostConfigArray, DNSResolverHostOverrideAlias{
		Host:   "test",
		Domain: "test.com",
		Descr:  "test"},
	)

	request := DNSResolverHostOverride{
		Host:    "test",
		Domain:  "test.com",
		IP:      []string{"0.0.0.0"},
		Aliases: DNSHostConfigArray}

	dnsResponse, error := client.DNS.CreateDNSResolverHostOverride(ctx, request)

	if error != nil {
		log.Fatalf("Error creating DNS resolver host override: %v", error)
	}

	fmt.Println("Succesfully Created DNSHostOverride")
	fmt.Println("Host Override ID Created", dnsResponse.Id)

	for _, alias := range dnsResponse.Aliases {
		fmt.Printf("Alias: %s\n", alias.Host)
	}

	dnsResponseDelete, error := client.DNS.DeleteDNSResolverHostOverride(ctx, request)

	fmt.Println("Deleted The following DNS ID Override ", dnsResponseDelete.Id)

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
