package csv

import (
	"encoding/csv"
	"io"
	"mm-db/internal/models"
	"os"
)

// ParseIPData reads IP data from a CSV file and returns a map of IP ranges to GeoData
func ParseIPData(filePath string) (map[string]models.GeoData, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the header (skip first line)
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// Create a map to store the IP data
	ipData := make(map[string]models.GeoData)

	// Read each row
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Check if the row has enough columns
		if len(row) < 7 {
			continue // Skip rows with insufficient data
		}

		// Create a GeoData struct for this row
		data := models.GeoData{
			GeoIsoCode:   row[1],
			State:        row[2],
			City:         row[3],
			IspName:      row[4],
			ConnType:     row[5],
			VpnProxyType: row[6],
		}

		// Add the data to the map with the IP range as the key
		ipData[row[0]] = data
	}

	return ipData, nil
}
