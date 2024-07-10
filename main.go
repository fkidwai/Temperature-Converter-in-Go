package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Global variables
var (
	originUnit         string  // Temperature unit to convert from
	originValue        float64 // Temperature value to convert
	shouldConvertAgain string  // User response to continue conversion
	err                error   // Generic error variable

	// Predefined errors
	errInvalidArguments = errors.New("Invalid arguments")
	errReadingInput     = errors.New("Error reading input")
)

func main() {
	// Check for valid command-line arguments
	if len(os.Args) != 2 {
		printError(errInvalidArguments)
	}

	// Get the temperature unit and convert to uppercase
	originUnit = strings.ToUpper(os.Args[1])

	for {
		// Prompt user for the temperature value
		fmt.Print("What is the current temperature in " + originUnit + " ? ")
		_, err := fmt.Scanln(&originValue)
		if err != nil {
			printError(errReadingInput)
		}

		// Convert the temperature based on the unit
		if originUnit == "C" {
			convertToFahrenheit(originValue)
		} else {
			convertToCelsius(originValue)
		}

		// Ask the user if they want to convert another temperature
		fmt.Print("Would you like to convert another temperature ? (y/n) ")
		_, err = fmt.Scanln(&shouldConvertAgain)
		if err != nil {
			printError(errReadingInput)
		}

		// Exit loop if the user doesn't want to convert another temperature
		if strings.ToUpper(strings.TrimSpace(shouldConvertAgain)) != "Y" {
			fmt.Println("Good bye!")
			break
		}
	}
}

// printError prints an error message and exits the program
func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

// convertToCelsius converts Fahrenheit to Celsius and prints the result
func convertToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%v F = %.0f C\n", value, convertedValue)
}

// convertToFahrenheit converts Celsius to Fahrenheit and prints the result
func convertToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%v C = %.0f F\n", value, convertedValue)
}
