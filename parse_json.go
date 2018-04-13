package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// BlacklistedResult is the actual data from various sources that tell whether
// this IP is blacklisted. Contains info from an individual source.
type BlacklistedResult struct {
	name        string
	link        string
	blacklisted bool
}

// Blacklisted is the subsection of the MullvadResult that shows whether you're
// blacklisted.
type Blacklisted struct {
	blacklisted bool
	results     [2]BlacklistedResult
}

// MullvadResult is the parent result of a query
type MullvadResult struct {
	ip                    string
	country               string
	city                  string
	longitude             float32
	latitude              float32
	mullvadExitIP         bool
	mullvadExitIPHostname string
	organization          string
	mullvadServerType     string
	blacklisted           Blacklisted
}

func main() {
	f, err := os.Open("sample.json")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	var data MullvadResult
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(data.blacklisted.results)
	fmt.Printf("%+v", data)
}
