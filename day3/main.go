package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strconv"
)

func isGear(num []int, lineNum int, lines [][]string) (int, int, bool) {

	firstNumIdx := num[0] - 1

	if firstNumIdx < 0 {
		firstNumIdx = 0
	}

	lastNumIdx := num[len(num)-1] + 1

	if lastNumIdx > len(lines[lineNum])-1 {
		lastNumIdx = len(lines[lineNum]) - 1
	}

	// first iterate over prev line, search for non-dot non-digit char
	if lineNum > 0 {
		for i, char := range lines[lineNum-1][firstNumIdx : lastNumIdx+1] {
			if char == "*" {
				return lineNum - 1, i + firstNumIdx, true
			}
		}
	}
	// then check prev and next char
	if lines[lineNum][firstNumIdx] == "*" {
		return lineNum, firstNumIdx, true
	}
	if lines[lineNum][lastNumIdx] == "*" {
		return lineNum, lastNumIdx, true
	}
	// then iterate over last line
	if lineNum < len(lines)-1 {
		for i, char := range lines[lineNum+1][firstNumIdx : lastNumIdx+1] {
			if char == "*" {
				return lineNum + 1, i + firstNumIdx, true
			}
		}
	}

	return 0, 0, false
}

func isPart(num []int, lineNum int, lines [][]string) bool {

	firstNumIdx := num[0] - 1

	if firstNumIdx < 0 {
		firstNumIdx = 0
	}

	lastNumIdx := num[len(num)-1] + 1

	if lastNumIdx > len(lines[lineNum])-1 {
		lastNumIdx = len(lines[lineNum]) - 1
	}

	// first iterate over prev line, search for non-dot non-digit char
	if lineNum > 0 {
		for _, char := range lines[lineNum-1][firstNumIdx : lastNumIdx+1] {
			if char != "." && !utils.IsDigit(char) {
				return true
			}
		}
	}
	// then check prev and next char
	if lines[lineNum][firstNumIdx] != "." && !utils.IsDigit(lines[lineNum][firstNumIdx]) {
		return true
	}
	if lines[lineNum][lastNumIdx] != "." && !utils.IsDigit(lines[lineNum][lastNumIdx]) {
		return true
	}
	// then iterate over last line
	if lineNum < len(lines)-1 {
		for _, char := range lines[lineNum+1][firstNumIdx : lastNumIdx+1] {
			if char != "." && !utils.IsDigit(char) {
				return true
			}
		}
	}

	return false
}

func buildNum(num []int, line []string) (int, error) {
	fullNum := ""

	for _, k := range num {
		fullNum += line[k]
	}

	// fmt.Println(fullNum)

	return strconv.Atoi(fullNum)
}

func main() {
	fmt.Println("part 1 answer:", part1())
	fmt.Println("part 2 answer:", part2())
}

func part1() int {
	lines := utils.LinesToMatrix(utils.InitAndReadFile())

	sum := 0

	for i, line := range lines {

		buildingNum := false
		num := make([]int, 0)

		for j, char := range line {
			if utils.IsDigit(char) {
				buildingNum = true
				num = append(num, j)
			} else {
				if buildingNum == true && isPart(num, i, lines) {
					numVal, err := buildNum(num, line)
					if err != nil {
						fmt.Println("Error building num:" + err.Error())
					}

					sum += numVal
				}

				buildingNum = false
				num = make([]int, 0)
			}
		}

		if buildingNum == true && isPart(num, i, lines) {
			numVal, err := buildNum(num, line)
			if err != nil {
				fmt.Println("Error building num:" + err.Error())
			}

			sum += numVal
		}
	}

	return sum
}

func addToGearMap(x int, y int, number int, gearMap map[string][]int) {
	key := strconv.Itoa(x) + "," + strconv.Itoa(y)

	if entry, exists := gearMap[key]; exists {
		gearMap[key] = append(entry, number)
	} else {
		gearMap[key] = []int{number}
	}
}

func part2() int {
	lines := utils.LinesToMatrix(utils.InitAndReadFile())

	sum := 0

	gearMap := make(map[string][]int, 0)

	for i, line := range lines {

		buildingNum := false
		num := make([]int, 0)

		for j, char := range line {
			if utils.IsDigit(char) {
				buildingNum = true
				num = append(num, j)
			} else {
				if buildingNum == true {
					if x, y, itIsAGear := isGear(num, i, lines); itIsAGear {
						numVal, err := buildNum(num, line)
						if err != nil {
							fmt.Println("Error building num:" + err.Error())
						}

						addToGearMap(x, y, numVal, gearMap)

					}
				}

				buildingNum = false
				num = make([]int, 0)
			}
		}

		if buildingNum == true {
			if x, y, itIsAGear := isGear(num, i, lines); itIsAGear {
				numVal, err := buildNum(num, line)
				if err != nil {
					fmt.Println("Error building num:" + err.Error())
				}

				addToGearMap(x, y, numVal, gearMap)

			}
		}
	}

	// process gearMap

	for key, value := range gearMap {
		fmt.Println(key, value)
		if len(value) == 2 {
			sum += value[0] * value[1]
		}
	}

	return sum
}
