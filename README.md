# Temperature Converter CLI

This is a simple Go program that converts temperatures between Celsius and Fahrenheit. The program prompts the user to input a temperature value and the unit (Celsius or Fahrenheit), then performs the conversion and displays the result. The user can continue converting temperatures until they choose to exit the program.

## Features

**Command-Line Argument:** The program accepts a single command-line argument to specify the temperature unit (either "C" for Celsius or "F" for Fahrenheit).

**User Input:** The program prompts the user to input the temperature value to be converted.

**Conversion Logic:**
Converts Celsius to Fahrenheit.
Converts Fahrenheit to Celsius.

**Interactive Loop:** The program continues to prompt the user for more conversions until the user decides to stop.

**Error Handling:** Handles invalid command-line arguments and input errors gracefully.

## Usage

Compile with `go build -o temp`.

Then, invoke the binary passing as argument the unit of temperature we want to convert **from**.
For example:

`./temp C` to convert from Celsius to Fahrenheit or `./temp F` to convert from Fahrenheit to Celsius.

The program will ask for the temperature value and perform the conversion.
After each conversion, it will prompt you if you want to convert another temperature.
Enter 'y' or 'Y' to continue, and 'n' or 'N' to exit the program.