package main

import (
	"strings"
	"regexp"
	"time"

)

func processInputFile(inputText string, airports []AirportInfo) string {
	//Handle IATA- and ICAO-codes
	processedText := convertAirportNames(inputText, airports)

	//Handle dates and times 
	processedText = convertDatesAndTimes(processedText)

	//Trim white space
	processedText = trimWhiteSpace(processedText)

	return processedText
}

func convertAirportNames(inputText string, airports []AirportInfo) string {
	for _, airport := range airports {
		// Regular expressions to match airport codes
		iataRegexCity := regexp.MustCompile(`\*#` + airport.IataCode + `\b`)
        icaoRegexCity := regexp.MustCompile(`\*##` + airport.IcaoCode + `\b`)
		iataRegex := regexp.MustCompile(`#` + airport.IataCode + `\b`)
		icaoRegex := regexp.MustCompile(`##` + airport.IcaoCode + `\b`)

		// Replace city codes
		inputText = iataRegexCity.ReplaceAllString(inputText, airport.Municipality)
		inputText = icaoRegexCity.ReplaceAllString(inputText, airport.Municipality)
		
		// Replace airport codes
		inputText = iataRegex.ReplaceAllString(inputText, airport.Name)
		inputText = icaoRegex.ReplaceAllString(inputText, airport.Name)
	}
	return inputText
}

func convertDatesAndTimes(inputText string) string {
    // Regular expression to match date/time formats
    dateTimePattern := regexp.MustCompile(`(D|T12|T24)\(([^)]+)\)`)
    
	// Replace date/time patterns with formatted versions
    processedText := dateTimePattern.ReplaceAllStringFunc(inputText, func(match string) string {
        // Find submatches in the current match
		parts := dateTimePattern.FindStringSubmatch(match)
        if len(parts) > 0 {
            switch parts[1] {
            case "D":
				if strings.HasSuffix(parts[2], "Z") {
					t, err := time.Parse("2006-01-02T15:04Z", parts[2])
					if err == nil {
						return t.Format("02 Jan 2006")
					}
				} else {
					t, err := time.Parse("2006-01-02T15:04-07:00", parts[2])
					if err == nil {
						return t.Format("02 Jan 2006")
					}
				}
            case "T12":
				if strings.HasSuffix(parts[2], "Z") {
					t, err := time.Parse("2006-01-02T15:04Z", parts[2])
					if err == nil {
						return t.Format("03:04PM (+00:00)")
					}
				} else {
					t, err := time.Parse("2006-01-02T15:04-07:00", parts[2])
					if err == nil {
						return t.Format("03:04PM (-07:00)")
					}
				}
            case "T24":
				if strings.HasSuffix(parts[2], "Z") {
					t, err := time.Parse("2006-01-02T15:04Z", parts[2])
					if err == nil {
						return t.Format("15:04 (+00:00)")
					}
				} else {
					t, err := time.Parse("2006-01-02T15:04-07:00", parts[2])
					if err == nil {
						return t.Format("15:04 (-07:00)")
					}
				}
            }
        }
        return match // Return unchanged if no valid conversion
    })
    return processedText
}

func trimWhiteSpace(inputText string) string {
	//Replace line-break charecters (\v, \f, \r) with a new-line character (\n)
	processedText := strings.ReplaceAll(inputText, "\v", "\n")
	processedText = strings.ReplaceAll(processedText, "\f", "\n")
	processedText = strings.ReplaceAll(processedText, "\r", "\n")
	
	//Remove more than two consecutive new-lines
	processedText = regexp.MustCompile(`\n{3,}`).ReplaceAllString(processedText, "\n\n")

	return processedText
}