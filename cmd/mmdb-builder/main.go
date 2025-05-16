package main

import (
	"log"
	"net"
	"os"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/inserter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
)

func main() {
	// Create a new MMDB writer with the "sample" database type
	writer, err := mmdbwriter.New(mmdbwriter.Options{
		DatabaseType: "Sample-Data",
		Description: map[string]string{
			"en": "Sample MMDB with custom data",
		},
		IPVersion:          6,
		RecordSize:         24,
		IncludeReservedNetworks: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Define and insert data for Company A's Engineering department
	_, engNet, err := net.ParseCIDR("192.168.1.0/24")
	if err != nil {
		log.Fatal(err)
	}
	engData := mmdbtype.Map{
		"Company.DeptName": mmdbtype.String("Engineering"),
		"Company.Environments": mmdbtype.Slice{
			mmdbtype.String("development"),
			mmdbtype.String("staging"),
			mmdbtype.String("production"),
		},
		"Company.Location": mmdbtype.String("Building A"),
		"Company.AccessLevel": mmdbtype.Uint64(3),
	}
	if err := writer.InsertFunc(engNet, inserter.TopLevelMergeWith(engData)); err != nil {
		log.Fatal(err)
	}

	// Define and insert data for Company A's Marketing department
	_, mktNet, err := net.ParseCIDR("192.168.2.0/24")
	if err != nil {
		log.Fatal(err)
	}
	mktData := mmdbtype.Map{
		"Company.DeptName": mmdbtype.String("Marketing"),
		"Company.Environments": mmdbtype.Slice{
			mmdbtype.String("development"),
			mmdbtype.String("staging"),
		},
		"Company.Location": mmdbtype.String("Building B"),
		"Company.AccessLevel": mmdbtype.Uint64(2),
	}
	if err := writer.InsertFunc(mktNet, inserter.TopLevelMergeWith(mktData)); err != nil {
		log.Fatal(err)
	}

	// Define and insert data for Company A's Finance department
	_, finNet, err := net.ParseCIDR("192.168.3.0/24")
	if err != nil {
		log.Fatal(err)
	}
	finData := mmdbtype.Map{
		"Company.DeptName": mmdbtype.String("Finance"),
		"Company.Environments": mmdbtype.Slice{
			mmdbtype.String("production"),
		},
		"Company.Location": mmdbtype.String("Building C"),
		"Company.AccessLevel": mmdbtype.Uint64(4),
	}
	if err := writer.InsertFunc(finNet, inserter.TopLevelMergeWith(finData)); err != nil {
		log.Fatal(err)
	}

	// Define and insert data for Company A's HR department
	_, hrNet, err := net.ParseCIDR("192.168.4.0/24")
	if err != nil {
		log.Fatal(err)
	}
	hrData := mmdbtype.Map{
		"Company.DeptName": mmdbtype.String("HR"),
		"Company.Environments": mmdbtype.Slice{
			mmdbtype.String("development"),
		},
		"Company.Location": mmdbtype.String("Building B"),
		"Company.AccessLevel": mmdbtype.Uint64(1),
	}
	if err := writer.InsertFunc(hrNet, inserter.TopLevelMergeWith(hrData)); err != nil {
		log.Fatal(err)
	}

	// Write the database to disk
	fh, err := os.Create("sample-data.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	_, err = writer.WriteTo(fh)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully created sample-data.mmdb")
}
