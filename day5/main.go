package main

import (
	"adventofcode2023/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part 1 answer:", part1())
	fmt.Println("part 2 answer:", part2())
}

func parseLines() ([]string, []string, []string, []string, []string, []string, []string, []string) {
	blocks := utils.InitAndReadFileToBlocks()

	seeds := strings.Split(blocks[0], " ")[1:]

	toSoil := strings.Split(blocks[1], "\r\n")[1:]
	toFert := strings.Split(blocks[2], "\r\n")[1:]
	toWater := strings.Split(blocks[3], "\r\n")[1:]
	toLight := strings.Split(blocks[4], "\r\n")[1:]
	toTemp := strings.Split(blocks[5], "\r\n")[1:]
	toHum := strings.Split(blocks[6], "\r\n")[1:]
	toLoc := strings.Split(blocks[7], "\r\n")[1:]

	return seeds, toSoil, toFert, toWater, toLight, toTemp, toHum, toLoc

}

func part1() int {
	seeds, toSoil, toFert, toWater, toLight, toTemp, toHum, toLoc := parseLines()

	nums := toInts(seeds)

	nums = convert(nums, mapToIntList(toSoil))
	nums = convert(nums, mapToIntList(toFert))
	nums = convert(nums, mapToIntList(toWater))
	nums = convert(nums, mapToIntList(toLight))
	nums = convert(nums, mapToIntList(toTemp))
	nums = convert(nums, mapToIntList(toHum))
	nums = convert(nums, mapToIntList(toLoc))

	min := math.MaxInt
	for _, n := range nums {
		if n < min {
			min = n
		}
	}

	return min
}

func convert(nums []int, maps [][]int) []int {

	result := make([]int, len(nums))

	for i, n := range nums {
	Nums:
		for j, m := range maps {
			sourceStart := m[1]
			sourceEnd := m[1] + m[2]

			if n >= sourceStart && n < sourceEnd {
				result[i] = (n - sourceStart) + m[0]
				break Nums
			} else if j == len(maps)-1 {
				// if not found in map
				result[i] = n
			}
		}
	}

	return result
}

func mapToIntList(convMap []string) [][]int {
	result := make([][]int, 0)
	for _, item := range convMap {
		result = append(result, toInts(strings.Split(item, " ")))
	}

	return result
}

func toInts(vals []string) []int {
	result := make([]int, 0)
	for _, val := range vals {
		i, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("error converting to int", val)
		}
		result = append(result, i)
	}

	return result
}

func part2() int {
	seeds, toSoil, toFert, toWater, toLight, toTemp, toHum, toLoc := parseLines()
	// seeds, toSoil, _, _, _, _, _, _ := parseLines()

	nums := convSeedRanges(toInts(seeds))
	fmt.Println("Got seeds")

	nums = filterNumber(nums, intListToBags(mapToIntList(toSoil)))
	fmt.Println("Got soil", nums)

	nums = filterNumber(nums, intListToBags(mapToIntList(toFert)))
	fmt.Println("Got fert", nums)

	nums = filterNumber(nums, intListToBags(mapToIntList(toWater)))
	fmt.Println("Got water", nums)

	nums = filterNumber(nums, intListToBags(mapToIntList(toLight)))
	fmt.Println("Got light", nums)

	nums = filterNumber(nums, intListToBags(mapToIntList(toTemp)))
	fmt.Println("Got temp", nums)

	nums = filterNumber(nums, intListToBags(mapToIntList(toHum)))
	fmt.Println("Got hum", nums)

	nums = filterNumber(nums, intListToBags(mapToIntList(toLoc)))
	fmt.Println("Got loc", nums)

	min := math.MaxInt
	for _, i := range nums {
		if i[0] > 0 && i[0] < min {
			min = i[0]
		}
	}

	return min
}

func convSeedRanges(seeds []int) [][]int {
	result := make([][]int, 0)

	for i := 0; i < len(seeds); i += 2 {
		result = append(result, []int{seeds[i], seeds[i+1]})
	}

	return result
}

type Bag struct {
	Dest   int
	Source int
	Range  int
}

func intListToBags(intList [][]int) []Bag {
	bags := make([]Bag, 0)

	for _, i := range intList {
		bags = append(bags, Bag{i[0], i[1], i[2]})
	}

	return bags
}

/*
*
bags is something like
10 50 5
10 = dest
50 = source
5 = range span

this returns a set of ranges in start, range order
[]int
10 5 12 7
*/
func filterNumber(ranges [][]int, bags []Bag) [][]int {
	newRanges := make([][]int, 0)

	// Sort bags by smallest source first
	sort.Slice(bags, func(i, j int) bool {
		return bags[i].Source < bags[j].Source
	})

	for _, r := range ranges {
		s := r[0]
		remainingRange := r[1]

		for _, bag := range bags {
			if remainingRange <= 0 {
				break
			}

			bagStart := bag.Source
			bagEnd := bag.Source + bag.Range
			diff := bag.Source - bag.Dest

			// Check if the current start is before the current bag
			if s < bagStart {
				overlap := intMin(bagStart-s, remainingRange)
				newRanges = append(newRanges, []int{s, overlap})
				s += overlap
				remainingRange -= overlap
			}

			// Check if there's still a range left to process within the current bag
			if remainingRange > 0 && s >= bagStart && s < bagEnd {
				overlap := intMin(bagEnd-s, remainingRange)
				newRanges = append(newRanges, []int{s - diff, overlap})
				s += overlap
				remainingRange -= overlap
			}
		}

		// Add any remaining range that wasn't covered by a bag
		if remainingRange > 0 {
			newRanges = append(newRanges, []int{s, remainingRange})
		}
	}

	return newRanges
}

func intMin(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
