package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eliucid8/advent-of-code-2025/internal/util"
)

func part1() string {
	lines, err := util.ReadLines("data/day01/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	displacement := 50
	sum := 0
	for _, l := range lines {
		if l == "" {
			continue
		}
		distance, _ := strconv.Atoi(l[1:])
		if l[0] == 'R' {
			displacement += distance
		} else {
			displacement -= distance
		}
		displacement %= 100
		if displacement < 0 {
			displacement += 100
		}
		// fmt.Println("displacement:", displacement, " distance:", distance)
		if displacement == 0 {
			sum += 1
		}
	}
	return strconv.Itoa(sum)
}

func part2() string {
	lines, err := util.ReadLines("data/day01/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	displacement := 50
	sum := 0
	for _, l := range lines {
		if l == "" {
			continue
		}
		distance, _ := strconv.Atoi(l[1:])
		// add 1 to distance for every 100 in the turn.
		sum += distance / 100
		// account for these extra turns
		distance %= 100
		if l[0] == 'R' {
			displacement += distance
		} else {
			// if the dial is pointing at 0, any move left causes it to go negative. We only want this to count if it passes 0 again.
			if displacement == 0 {
				displacement = 100
			}
			displacement -= distance
		}
		// if we crossed 0 again
		if displacement <= 0 || displacement >= 100 {
			sum += 1
		}
		// fmt.Println("displacement:", displacement, " distance:", distance, " sum:", sum)

		displacement %= 100
		if displacement < 0 {
			displacement += 100
		}
	}
	return strconv.Itoa(sum)
}

func main() {
	answer1 := part1()
	fmt.Println("part 1:", answer1)
	answer2 := part2()
	fmt.Println("part 2:", answer2)
}
