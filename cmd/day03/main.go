package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eliucid8/advent-of-code-2025/internal/util"
)

func part1(lines []string) string {
	total := 0
	for _, line := range lines {
		maxfirst, maxfirstIndex := 0, 0
		maxlast := 0
		for i, char := range line[:len(line)-1] {
			var digit = int(char - '0')
			if digit > maxfirst {
				maxfirst = digit
				maxfirstIndex = i
			}
		}
		for _, char := range line[maxfirstIndex+1:] {
			var digit = int(char - '0')
			if digit > maxlast {
				maxlast = digit
			}
		}
		total += maxfirst*10 + maxlast
	}

	return strconv.Itoa(total)
}

func part2(lines []string) string {
	var total uint64 = 0
	const numDigits = 12
	for _, line := range lines {
		var jValue uint64 = 0
		bankIdx := 0
		for jIdx := 0; jIdx < numDigits; jIdx++ {
			maxDigit := 0
			selectionIdx := 0
			for i, char := range line[bankIdx : len(line)-11+jIdx] {
				var digit = int(char - '0')
				if digit > maxDigit {
					maxDigit = digit
					selectionIdx = i
				}
			}
			bankIdx += selectionIdx + 1
			jValue *= 10
			jValue += uint64(maxDigit)
		}
		total += jValue
		// fmt.Println(jValue)
	}
	return strconv.FormatUint(total, 10)
}

func main() {
	lines, err := util.ReadLines("data/day03/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	answer1 := part1(lines)
	fmt.Println("part 1:", answer1)
	answer2 := part2(lines)
	fmt.Println("part 2:", answer2)
}
