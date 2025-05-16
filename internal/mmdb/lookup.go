package mmdb

import (
	"mm-db/internal/models"
	"net"

	"github.com/oschwald/maxminddb-golang"
)

// Lookup service for retrieving IP data from an MMDB file
type Lookup struct {
	reader *maxminddb.Reader
}

// NewLookup creates a new Lookup service
func NewLookup(mmdbPath string) (*Lookup, error) {
	reader, err := maxminddb.Open(mmdbPath)
	if err != nil {
		return nil, err
	}

	return &Lookup{
		reader: reader,
	}, nil
}

// GetIPData retrieves geo data for an IP address
func (l *Lookup) GetIPData(ipStr string) (*models.GeoData, error) {
	// Parse the IP address
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, nil
	}

	// Look up the IP in the MMDB
	var data models.GeoData
	err := l.reader.Lookup(ip, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Close closes the MMDB reader
func (l *Lookup) Close() error {
	return l.reader.Close()
}
