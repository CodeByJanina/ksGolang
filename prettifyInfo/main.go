package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type AirportInfo struct {
	Name string
	IsoCountry string
	Municipality string
	IcaoCode string
	IataCode string
	Coordinates string
}

func main() {

	//Define boolean flags
	helpFlag := flag.Bool("h", false, "Display usage information\nUsage: $ go run . -h")
	formatFlag := flag.Bool("f", false, "Print formated output to screen\nUsage: $ go run . -f ./input.txt ./output.txt ./airport-lookup.csv")
	dynamicColumnsFlag := flag.Bool("d", false, "Use dynamic column order\nUsage: $ go run . -d ./input.txt ./output.txt ./airport-lookup.csv")
	findAirportFlag := flag.Bool("i", false, "Find airport's information\nUsage: $ go run . -i")
	
	//Parse the command-line flags
	flag.Parse()

	if *findAirportFlag {
		findAirportInfo()
		return
	}

	if *helpFlag || flag.NArg() != 3 {
        printHelp()
        return
    }

	//Extract command-line argument
	inputFile := os.Args[len(os.Args)-3]
	outputFile := os.Args[len(os.Args)-2]
	airportLookupFile := os.Args[len(os.Args)-1]

	// Check if the input or airport lookup file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		log.Fatal("Input not found")
	}

	if _, err := os.Stat(airportLookupFile); os.IsNotExist(err) {
		log.Fatal("Airport lookup not found")
	}

	// Check if the airport lookup data is malformed
	if err := checkMalformedAirportLookup(airportLookupFile); err != nil {
		log.Fatal("Airport lookup malformed:", err)
	}

	var airports []AirportInfo
    var err error
	
    if *dynamicColumnsFlag {
        airports, err = readAirportLookupWithDynamicColumns(airportLookupFile)
    } else {
        airports, err = readAirportLookup(airportLookupFile)
    }
    if err != nil {
        log.Fatal("Error reading airport lookup:", err)
    }

	// Read the input file
	inputText, err := readInputFile(inputFile)
	if err != nil {
		log.Fatal("Error reding input file", err)
	}

	// Process the itinerary (replace codes with names, handle dates/times, trim whitespace)
	processedText := processInputFile(inputText, airports)

	//Write processed text to ouput file.
	err = writeOutputFile(outputFile, processedText)
	if err != nil {
		log.Fatal("Error writing output file:", err)
	}	

	// If the format flag is set, print formatted output to stdout
	if *formatFlag {
		formattedOutput := formatOutput(processedText, airports)
		fmt.Print("\n", formattedOutput, "\n")
	} 
}

func printHelp() {
	fmt.Println("Itinerary usage:")
	fmt.Println("$ go run . [options] ./input.txt ./output.txt ./airport-lookup.csv")
	fmt.Println("Options:")
	flag.PrintDefaults()
}


