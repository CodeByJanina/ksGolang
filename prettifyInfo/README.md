# Itinerary Prettifier

This is a command line tool which can prettify flight itineraries. The tool reads an itinerary from a text file(input), processes the text making it customer-friendly, and writes the result to a new file (output). It uses airport lookup CSV to extract airport names or city names for given airport code. Dates and times will be changed from ISO 8601 standard to customer friendly dates and times. Also extra white spaces will be trimmed.

## Learning outcomes

- String manipulation
- Command line arguments
- Navigating the file system
- Reading from files
- Writing to files

## Usage

The tool is run from the command line with three arguments:
    1. Path to the input file
    2. Path to the output file
    3. Path to the airport lookup file

    $ go run . ./input.txt ./output.txt ./airport-lookup.csv

Running a -h flag will display the usage:    

    $ go run . -h
    Itinerary usage:
    $ go run . [options] ./input.txt ./output.txt ./airport-lookup.csv
    Options:
      -d    Use dynamic column order
            Usage: $ go run . -d ./input.txt ./output.txt ./airport-lookup.csv
      -f    Print formated output to screen
            Usage: $ go run . -f ./input.txt ./output.txt ./airport-lookup.csv
      -h    Display usage information
            Usage: $ go run . -h
      -i    Find airport's information
            Usage: $ go run . -i

The program has optional flags which are displayed in the usage.

### Optional flags
* `-d` The program can be used when columns in the airport lookup file have been reordered.
* `-f` Formats the output with colors, underlines and bolding. The formated output is printed to the screen.
* `-h` Shows the usage information with short descriptions of optional flags.
* `-i` User provides airport lookup's file name then the airport's name, ICAO or IATA code and the tool will find airport's         information and prints it to the screen. 

## Text processing
To make the itinerary more customer friendly selected data will be processed. For example:

### Airport names
* IATA code: #LAX to Los Angeles International Airport
* ICAO code: ##EGLL to London Heathrow Airport

### City names
* IATA code: *#LHR to London
* ICAO code: *##KLAX to Los Angeles

### Dates and times
* Dates: D(2007-04-05T12:30-02:00) to 05 Apr 2007
* 12 Hour time: T12(2007-04-05T12:30-02:00) to 12:30PM(-02:00)
* 24 Hour time: T24(2007-04-05T12:30-02:00) to 12:30(-02:00)

 If the offset is Z (zulu-time) it will be changed to (+00:00)

 ### White space
 * Line-break characters (`\v`, `\f`, `\r`) are converted to a new-line character: `\n`. 
 * If there is more than one consecutive blank line, it will be changed to one blank line.
