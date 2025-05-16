package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"mm-db/internal/mmdb"
	"os"
)

func main() {
	// Parse command line arguments
	mmdbFile := flag.String("db", "custom-geo-ipblocks.mmdb", "Path to the MMDB file")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: mmdb-lookup -db=<path_to_mmdb> <ip_address> [<ip_address>...]")
		os.Exit(1)
	}

	// Create a new lookup service
	lookup, err := mmdb.NewLookup(*mmdbFile)
	if err != nil {
		log.Fatalf("Failed to open MMDB file: %v", err)
	}
	defer lookup.Close()

	// Look up each IP address
	for _, ipStr := range flag.Args() {
		fmt.Printf("Looking up IP: %s\n", ipStr)
		
		data, err := lookup.GetIPData(ipStr)
		if err != nil {
			fmt.Printf("Error looking up IP %s: %v\n", ipStr, err)
			continue
		}

		if data == nil {
			fmt.Printf("No data found for IP: %s\n", ipStr)
			continue
		}

		// Pretty print the result
		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Printf("Error formatting result: %v\n", err)
			continue
		}

		fmt.Println(string(jsonData))
		fmt.Println()
	}
}
