package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/eliucid8/advent-of-code-2025/internal/util"
)

func createRanges(lines []string) [][2]uint64 {
	var rawranges = strings.Split(lines[0], ",")
	var ranges [][2]uint64
	for _, rawrange := range rawranges {
		var ends = strings.Split(rawrange, "-")
		var currange [2]uint64
		currange[0], _ = strconv.ParseUint(ends[0], 10, 64)
		currange[1], _ = strconv.ParseUint(ends[1], 10, 64)
		ranges = append(ranges, currange)
	}
	return ranges
}

func part1(lines []string) string {
	invalidRanges := createRanges(lines)

	var invalidSum uint64

	for _, invalidRange := range invalidRanges {

		// takes the number of digits of top of range, halves, + 1, gets base mult factor
		// e.g. 9876 -> 0101
		var numDigits = int(math.Floor(math.Log10(float64(invalidRange[1])))) + 1
		var baseMult uint64 = uint64(math.Pow10(numDigits/2)) + 1
		// fmt.Printf("bM: %d ", baseMult)
		// max value for repetition is 10^(nD // 2) - 1, e.g. 99
		// min value is 10^((nD // 2) - 1), e.g. 10
		var top = min(invalidRange[1]/baseMult, uint64(math.Pow10(numDigits/2))-1)
		// round up
		var bot = max((invalidRange[0]+baseMult-1)/baseMult, uint64(math.Pow10(numDigits/2-1)))
		// fmt.Printf("b: %d t: %d\n", bot, top)
		if bot > top {
			continue
		}
		invalidSum += ((bot + top) * (top - bot + 1) / 2) * baseMult
	}

	return strconv.FormatUint(invalidSum, 10)
}

func part2(lines []string) string {
	ranges := createRanges(lines)

	var invalidSum uint64
	var invalidSet = make(map[uint64]bool)

	for _, invalidRange := range ranges {
		// fmt.Printf("[%d-%d]\n", invalidRange[0], invalidRange[1])
		var numDigitsLow = int(math.Floor(math.Log10(float64(invalidRange[0])))) + 1
		var numDigitsHigh = int(math.Floor(math.Log10(float64(invalidRange[1])))) + 1
		for spacing := 1; spacing <= numDigitsHigh/2; spacing++ {
			// that's real tricky.
			// single digits don't count as a repeat: "2" is valid
			// fortunately that was the last testcase :)
			if numDigitsLow != 1 {
				invalidSum = trySpacing(numDigitsLow, spacing, invalidRange, invalidSum, invalidSet)
			}

			if numDigitsLow == numDigitsHigh {
				continue
			}

			invalidSum = trySpacing(numDigitsHigh, spacing, invalidRange, invalidSum, invalidSet)
		}
		// fmt.Print("\n")
	}
	return strconv.FormatUint(invalidSum, 10)
}

func trySpacing(numDigitsLow int, spacing int, invalidRange [2]uint64, invalidSum uint64, invalidSet map[uint64]bool) uint64 {
	var baseMult uint64 = 1
	for i := 0; i < numDigitsLow-spacing; i += spacing {
		baseMult *= uint64(math.Pow10(spacing))
		baseMult += 1
	}
	// fmt.Printf("bM: %d ", baseMult)
	var top = min(invalidRange[1]/baseMult, uint64(math.Pow10(spacing))-1)
	var bot = max((invalidRange[0]+baseMult-1)/baseMult, uint64(math.Pow10(spacing-1)))
	if bot <= top {
		// fmt.Printf("b: %d t: %d | ", bot, top)
		for i := bot; i <= top; i++ {
			if !invalidSet[i*baseMult] {
				invalidSet[i*baseMult] = true
				invalidSum += i * baseMult
			}
		}
	}
	return invalidSum
}

func main() {
	lines, err := util.ReadLines("data/day02/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	answer1 := part1(lines)
	fmt.Println("part 1:", answer1)
	answer2 := part2(lines)
	fmt.Println("part 2:", answer2)
}
