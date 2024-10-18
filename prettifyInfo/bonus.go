package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"regexp"
	"strings"
    "fmt"
    "log"
)

func formatOutput(output string, airports []AirportInfo) string {
	const (
		Blue   = "\033[34m"
		Yellow = "\033[33m"
		Bold = "\033[1m"
		Underline = "\033[4m"
		Reset  = "\033[0m"
	)

	// Define regular expressions for matching dates, 12-hour times, and 24-hour times
	dateRegex := regexp.MustCompile(`\d{2} \w{3} \d{4}`)	
	time12HourRegex := regexp.MustCompile(`\d{2}:\d{2}(?:AM|PM)\s\([+-]\d{2}:\d{2}\)`)
	time24HourRegex := regexp.MustCompile(`\d{2}:\d{2}\s\([+-]\d{2}:\d{2}\)`)
    offsetRegex := regexp.MustCompile(`\([+-]\d{2}:\d{2}\)`)

	for _, airport := range airports {
		output = strings.ReplaceAll(output, airport.Name, Blue+airport.Name+Reset)
	}

	output = dateRegex.ReplaceAllStringFunc(output, func(match string) string {
		return Underline + match + Reset
	})

	output = time12HourRegex.ReplaceAllStringFunc(output, func(match string) string {
		return Yellow + match + Reset
	})

	output = time24HourRegex.ReplaceAllStringFunc(output, func(match string) string {
		return Yellow + match + Reset
	})

    output = offsetRegex.ReplaceAllStringFunc(output, func(match string) string {
		return Bold + match + Reset
	})   

	return output
}

func readAirportLookupWithDynamicColumns(airportLookupFile string) ([]AirportInfo, error) {
    file, err := os.Open(airportLookupFile)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(bufio.NewReader(file))

    // Read the header row to determine column positions
    header, err := reader.Read()
    if err != nil {
        return nil, err
    }

    // Map column names to their positions
    columnPositions := make(map[string]int)
    for i, columnName := range header {
        columnPositions[columnName] = i
    }

    // Initialize slice to hold AirportInfo objects
    var airports []AirportInfo

    // Read remaining rows
    for {
        record, err := reader.Read()
        if record == nil {
            break // Exit loop if there are no more records to read
        }
        if err != nil {
            return nil, err
        }

        // Create AirportInfo object using dynamically determined column positions
        airport := AirportInfo{
            Name:        record[columnPositions["name"]],
            IsoCountry:  record[columnPositions["iso_country"]],
            Municipality: record[columnPositions["municipality"]],
            IcaoCode:    record[columnPositions["icao_code"]],
            IataCode:    record[columnPositions["iata_code"]],
            Coordinates: record[columnPositions["coordinates"]],
        }
        airports = append(airports, airport)
    }
    return airports, nil
}

func findAirportInfo() {
    var airportLookupFile string
    var airports []AirportInfo
    var err error
    
    fmt.Println("Find airport's information.")
    fmt.Println("Give airport lookup file name:")
    fmt.Scanln(&airportLookupFile)

    // Check if the airport lookup data is malformed
	if err := checkMalformedAirportLookup(airportLookupFile); err != nil {
		log.Fatal("Airport lookup malformed:", err)
	}

    airports, err = readAirportLookup(airportLookupFile)
    if err != nil {
        log.Fatal("Error reading airport lookup:", err)
    }

    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Give the airport's name, ICAO-code or IATA-code:")
    // Read the user input until newline character
    query, _ := reader.ReadString('\n')

    // Remove leading and trailing whitespaces from the input
    query = strings.TrimSpace(query)

    var found bool

    for i, airport := range airports {
        if airport.Name == query || airport.IcaoCode == query || airport.IataCode == query {
            fmt.Println("\nFound Airport:", airport.Name)
            fmt.Println("Municipality:", airport.Municipality)
            fmt.Println("ISO Country:", airport.IsoCountry)
            fmt.Println("ICAO Code:", airport.IcaoCode)
            fmt.Println("IATA Code:", airport.IataCode)
            fmt.Println("Coordinates:", airport.Coordinates)
            fmt.Println("Row Number:", i+1) 
            found = true
            break
        }
    }
    
    if !found {
        fmt.Println("Airport not found from airport lookup:", query)
    }
}
