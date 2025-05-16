package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/oschwald/maxminddb-golang"
)

func main() {
	// Parse command line arguments
	dbFile := flag.String("db", "sample-data.mmdb", "Path to the MMDB file")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: mmdb-query -db=<path_to_mmdb> <ip_address>")
		os.Exit(1)
	}

	// Open the MMDB file
	db, err := maxminddb.Open(*dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Print database metadata
	fmt.Printf("Database Information:\n")
	fmt.Printf("  Description: %v\n", db.Metadata.Description)
	fmt.Printf("  IP Version: %d\n", db.Metadata.IPVersion)
	fmt.Printf("  Database Type: %s\n\n", db.Metadata.DatabaseType)

	// Query each IP address provided
	for _, ipStr := range flag.Args() {
		ip := net.ParseIP(ipStr)
		if ip == nil {
			fmt.Printf("Invalid IP address: %s\n", ipStr)
			continue
		}

		var result map[string]interface{}
		err = db.Lookup(ip, &result)
		if err != nil {
			fmt.Printf("Error looking up %s: %s\n", ipStr, err)
			continue
		}

		// Pretty print the result
		fmt.Printf("Results for IP: %s\n", ipStr)
		if len(result) == 0 {
			fmt.Println("  No data found")
		} else {
			jsonData, err := json.MarshalIndent(result, "  ", "  ")
			if err != nil {
				fmt.Printf("  Error formatting result: %s\n", err)
			} else {
				fmt.Println(string(jsonData))
			}
		}
		fmt.Println()
	}
}
