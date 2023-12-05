package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("part 1 answer:", part1())
	fmt.Println("part 2 answer:", part2())
}

func part1() int {
	lines := utils.InitAndReadFile()

	sum := 0

	for _, line := range lines {
		a := strings.Split(line, ": ")
		b := strings.Split(a[1], " | ")

		winning := sliceToSet(strings.Split(b[0], " "))
		nums := strings.Split(b[1], " ")

		points := 0

		for _, num := range nums {
			if _, contains := winning[num]; num != "" && contains {
				if points == 0 {
					points = 1
				} else {
					points = points * 2
				}
			}
		}

		sum += points
		// fmt.Println(a[0], points)
	}

	return sum
}

func sliceToSet(slice []string) map[string]bool {
	set := make(map[string]bool)
	for _, item := range slice {
		if item != "" {
			set[item] = true
		}
	}
	return set
}

func part2() int {

	lines := utils.InitAndReadFile()
	cardCount := len(lines)

	ops := 0

	pointLookup, queue := calcTotalsAndAddToQueue(lines)

	for len(queue) > 0 {
		// pop off queue
		item := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		for i := item + 1; i < item+pointLookup[item]+1 && i < cardCount; i++ {
			queue = append(queue, i)
			ops++
		}

	}

	return cardCount + ops
}

func calcTotalsAndAddToQueue(lines []string) (map[int]int, []int) {
	lookup := make(map[int]int)
	queue := make([]int, 0)

	for i, line := range lines {
		queue = append(queue, i)
		a := strings.Split(line, ": ")
		b := strings.Split(a[1], " | ")

		winning := sliceToSet(strings.Split(b[0], " "))
		nums := strings.Split(b[1], " ")

		points := 0

		for _, num := range nums {
			if _, contains := winning[num]; num != "" && contains {
				points++
			}
		}

		lookup[i] = points

	}

	return lookup, queue

}
