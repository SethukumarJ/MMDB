package mmdb

import (
	"log"
	"mm-db/internal/models"
	"net"
	"os"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/inserter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
)

// BuildMMDB creates an MMDB file from a map of IP ranges to GeoData
func BuildMMDB(ipData map[string]models.GeoData, outputPath string) error {
	// Create a new MMDB writer
	writer, err := mmdbwriter.New(mmdbwriter.Options{
		DatabaseType: "GeoIP-Custom",
		Description: map[string]string{
			"en": "Custom IP Geolocation Database",
		},
		IPVersion:              6,
		RecordSize:             24,
		IncludeReservedNetworks: true,
	})
	if err != nil {
		return err
	}

	// Insert each IP range and its data into the MMDB
	for ipRange, data := range ipData {
		// Parse the IP range
		_, network, err := net.ParseCIDR(ipRange)
		if err != nil {
			log.Printf("Error parsing IP range %s: %v", ipRange, err)
			continue
		}

		// Convert the GeoData to an mmdbtype.Map
		mmdbData := mmdbtype.Map{
			"geo_iso_code_2": mmdbtype.String(data.GeoIsoCode),
			"geo_state":      mmdbtype.String(data.State),
			"geo_city":       mmdbtype.String(data.City),
			"isp_name":       mmdbtype.String(data.IspName),
			"conn_type":      mmdbtype.String(data.ConnType),
			"vpn_proxy_type": mmdbtype.String(data.VpnProxyType),
		}

		// Insert the data into the MMDB
		if err := writer.InsertFunc(network, inserter.TopLevelMergeWith(mmdbData)); err != nil {
			log.Printf("Error inserting data for IP range %s: %v", ipRange, err)
			continue
		}
	}

	// Write the MMDB to a file
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = writer.WriteTo(file)
	return err
}
