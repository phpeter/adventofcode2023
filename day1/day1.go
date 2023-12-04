package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func checkForWord(line []string, words [][]string) (string, bool) {
	// words input will be something like {{"two","2"}, {"three","3"}}

	for _, entry := range words {
		// entry = ["two", "2"]
		if len(line) >= len(entry[0]) && strings.Join(line[:len(entry[0])], "") == entry[0] {
			return entry[1], true
		}
	}

	return "", false
}

var ltrs = map[string][][]string{
	"o": {{"one", "1"}},
	"t": {{"two", "2"}, {"three", "3"}},
	"f": {{"four", "4"}, {"five", "5"}},
	"s": {{"six", "6"}, {"seven", "7"}},
	"e": {{"eight", "8"}},
	"n": {{"nine", "9"}},
}

func findLastNumber(line []string) (string, bool) {
	for i := len(line) - 1; i > -1; i-- {
		if utils.IsDigit(line[i]) {
			return line[i], true
		}

		if entry, exists := ltrs[line[i]]; exists {
			if value, isThere := checkForWord(line[i:], entry); isThere {
				return value, true
			}
		}
	}

	return "", false
}

func findFirstNumber(line []string) (string, bool) {
	for i := 0; i < len(line); i++ {
		if utils.IsDigit(line[i]) {
			return line[i], true
		}

		if entry, exists := ltrs[line[i]]; exists {
			if value, isThere := checkForWord(line[i:], entry); isThere {
				return value, true
			}
		}
	}

	return "", false
}

func main() {
	lines := utils.InitAndReadFile()

	sum := 0

	for _, line := range lines {
		val := 0
		var err error

		chars := strings.Split(line, "")

		first, foundF := findFirstNumber(chars)
		last, foundL := findLastNumber(chars)

		if !foundF || !foundL {
			fmt.Println("error, could not find first or last number")
		} else {
			val, err = strconv.Atoi(first + last)

			if err != nil {
				fmt.Println("error, bad values for conversion: ", first, last)
			}

			fmt.Println("result:", val)

		}

		// fmt.Println(val)
		sum = sum + val
	}

	fmt.Println("Sum is:", sum)

}
