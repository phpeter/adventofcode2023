package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func LinesToMatrix(lines []string) [][]string {
	splitLines := make([][]string, len(lines))

	for i, line := range lines {
		splitLines[i] = strings.Split(line, "")
	}

	return splitLines
}

func IsDigit(char string) bool {
	_, err := strconv.Atoi(char)
	// err == nil means this is a digit
	// err != nil means this is a char
	return err == nil
}
