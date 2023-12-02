package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := utils.InitAndReadFile()

	redMax := 12
	greenMax := 13
	blueMax := 14

	sum := 0

	for _, line := range lines {
		semi := strings.Index(line, ":")
		gameId, _ := strconv.Atoi(line[5:semi])

		rounds := strings.Split(line[semi+2:], "; ")

		// fmt.Println("===GAME ", gameId, "===")

		gameSuccess := true

	Game:
		for _, round := range rounds {
			cubeCounts := strings.Split(round, ", ")

			for _, cubeCount := range cubeCounts {
				item := strings.Split(cubeCount, " ")
				count, _ := strconv.Atoi(item[0])
				// fmt.Println(cubeCount)
				switch item[1] {
				case "red":
					if count > redMax {
						gameSuccess = false
						break Game
					}
					break
				case "green":
					if count > greenMax {
						gameSuccess = false
						break Game
					}
					break
				case "blue":
					if count > blueMax {
						gameSuccess = false
						break Game
					}
					break
				}
			}
		}

		if gameSuccess {
			sum += gameId
		}
	}

	fmt.Println("part 1 result:", sum)
}

func part2() {
	lines := utils.InitAndReadFile()

	sum := 0

	for _, line := range lines {
		semi := strings.Index(line, ":")

		rounds := strings.Split(line[semi+2:], "; ")

		fmt.Println(line)
		redMax := 0
		greenMax := 0
		blueMax := 0

		for _, round := range rounds {
			cubeCounts := strings.Split(round, ", ")

			for _, cubeCount := range cubeCounts {
				item := strings.Split(cubeCount, " ")
				count, _ := strconv.Atoi(item[0])
				// fmt.Println(cubeCount)
				switch item[1] {
				case "red":
					if count > redMax {
						redMax = count
					}
					break
				case "green":
					if count > greenMax {
						greenMax = count
					}
					break
				case "blue":
					if count > blueMax {
						blueMax = count
					}
					break
				}
			}
		}

		fmt.Println("redMax:", redMax, "greenMax:", greenMax, "blueMax:", blueMax)

		sum += redMax * greenMax * blueMax
	}

	fmt.Println("part 2 result:", sum)
}
