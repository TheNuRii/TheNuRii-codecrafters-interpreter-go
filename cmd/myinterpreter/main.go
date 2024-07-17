package main 

import {
"fms"
"os"
}

const (
	LEFT_PAREN rune = '('
	RIGHT_PAREN rune = ')'
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}
	
	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintln(os.Stderr, "Unkonown command %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	rawfileContents, err := os.ReadFile(filename)
	if err !=  nil {
		fmt.Fprintln(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		panic("Scanner not implemented")
	} else {
		fmt.Println("E0F null")
	}

	fileContents := string(rawfileContents)
	
	for _, current := range fileContents {
		switch current {
		case LEFT_PAREN:
			fmt.Println("LEFT_PAREN ( null")

		case: RIGHT_PAREN:
			fmt.Println("RIGHT_PAREN ) null")
		}
	}

	fmt.Println("E0F null")
}