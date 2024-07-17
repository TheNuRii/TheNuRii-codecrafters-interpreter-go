package main
import (
	"fmt"
	"os"
)
func tokenize(fileContents []byte) {
	for _, c := range fileContents {
		switch c {
		case '(':
			fmt.Println("LEFT_PAREN ( null")
		case ')':
			fmt.Println("RIGHT_PAREN ) null")
		default:
			fmt.Fprintln(os.Stderr, "Unknown character:", c)
		}
	}
	// This must be the last line
	fmt.Println("EOF  null")
}
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}
	command := os.Args[1]
	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}
	// Uncomment this block to pass the first stage
	//
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	if len(fileContents) > 0 {
		panic("Scanner not implemented")
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
	tokenize(fileContents)
}