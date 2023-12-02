package utils

import (
	"bufio"
	"fmt"
	"os"
)

func InitAndReadFile() []string {
	// Check if there is exactly one command-line argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	// Get the input file path from the command-line argument
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Could not read file:", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Could not read file:", err.Error())
		os.Exit(1)
	}

	return lines
}
