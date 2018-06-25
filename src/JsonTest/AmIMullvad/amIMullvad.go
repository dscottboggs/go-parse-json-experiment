package AmIMullvad

import (
	"encoding/json"
	"log"
	"os"
)

// BlacklistedResult is the actual data from various sources that tell whether
// this IP is blacklisted. Contains info from an individual source.
type BlacklistedResult struct {
	Name        string
	Link        string
	Blacklisted bool
}

// Blacklisted is the subsection of the MullvadResult that shows whether you're
// blacklisted.
type Blacklisted struct {
	Blacklisted bool
	Results     map[string][2]BlacklistedResult
}

// MullvadResult is the parent result of a query
type MullvadResult struct {
	IP                    string
	Country               string
	City                  string
	Longitude             float32
	Latitude              float32
	MullvadExitIP         bool
	MullvadExitIPHostname string
	Organization          string
	MullvadServerType     string
	Blacklisted           map[string]Blacklisted
}

// ParseSample --
// Parses the sample JSON file.
func ParseSample() MullvadResult {
	var data MullvadResult
	f, err := os.Open("sample.json")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		log.Panic(err)
	}
	return data
}
