package main

import (
	"os"
	"bufio"
	"encoding/csv"
	"fmt"
)

func readInputFile(inputFile string) (string, error) {
	//Read the entire file content
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func checkMalformedAirportLookup(airportLookupFile string) error {
	//Open the airport lookup file
	file, err := os.Open(airportLookupFile)
	if err != nil {
		return err
	}
	//Defer closing the file until function exits
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = 6

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	//Iterate over each record in the CSV file 
	for _, record := range records {
		if len(record) != 6 {
			return fmt.Errorf("wrong number of fields")
		}
		for _, field := range record {
			if field == "" {
				return fmt.Errorf("missing or blank field found")
			}
		}
	}
	return nil
}

func readAirportLookup(airportLookupFile string) ([]AirportInfo, error) {
	file, err := os.Open(airportLookupFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = 6

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	//A slice to hold airport information
	var airports []AirportInfo

	for _, record := range records {
		// Create an AirportInfo struct with data from the record
		airport := AirportInfo{
			Name: record[0],
			IsoCountry: record[1],
			Municipality: record[2],
			IcaoCode: record[3],
			IataCode: record[4],
			Coordinates: record[5],
		}
		airports = append(airports, airport)
	}
	return airports, nil
}

func writeOutputFile(outputFile, content string) error {
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close() 

	_, err = output.WriteString(content)
	if err != nil {
		return err
	}

	fmt.Printf("%s created\n", outputFile)

	return nil
}
