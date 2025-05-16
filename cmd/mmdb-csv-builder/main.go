package main

import (
	"flag"
	"fmt"
	"log"
	"mm-db/internal/csv"
	"mm-db/internal/mmdb"
	"os"
	"path/filepath"
)

func main() {
	// Parse command line arguments
	csvFile := flag.String("csv", "data/ip_data.csv", "Path to the CSV file containing IP data")
	outputFile := flag.String("output", "custom-geo-ipblocks.mmdb", "Path to the output MMDB file")
	flag.Parse()

	// Ensure the CSV file exists
	if _, err := os.Stat(*csvFile); os.IsNotExist(err) {
		log.Fatalf("CSV file does not exist: %s", *csvFile)
	}

	// Create the output directory if it doesn't exist
	outputDir := filepath.Dir(*outputFile)
	if outputDir != "." && outputDir != "" {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			log.Fatalf("Failed to create output directory: %v", err)
		}
	}

	// Parse the CSV file
	fmt.Println("Parsing CSV file:", *csvFile)
	ipData, err := csv.ParseIPData(*csvFile)
	if err != nil {
		log.Fatalf("Failed to parse CSV file: %v", err)
	}
	fmt.Printf("Parsed %d IP ranges from CSV\n", len(ipData))

	// Build the MMDB file
	fmt.Println("Building MMDB file:", *outputFile)
	if err := mmdb.BuildMMDB(ipData, *outputFile); err != nil {
		log.Fatalf("Failed to build MMDB file: %v", err)
	}

	fmt.Println("Successfully created MMDB file:", *outputFile)
}
