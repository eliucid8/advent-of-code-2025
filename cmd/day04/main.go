package main

import (
	"fmt"
	"github.com/eliucid8/advent-of-code-2025/internal/util"
	"log"
	"strconv"
)

func part1() string {
	lines, err := util.ReadLines("data/day04/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	var rolls [][]bool

	for _, line := range lines {
		var row []bool
		for _, c := range line {
			row = append(row, c == '@')
		}
		rolls = append(rolls, row)
	}

	var directions [8][2]int = [8][2]int{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}

	count := 0
	for y, row := range rolls {
		for x, val := range row {
			if val {
				neighbors := 0
				for _, direction := range directions {
					if check_cell(x+direction[0], y+direction[1], rolls) {
						neighbors += 1
					}
				}
				if neighbors < 4 {
					count += 1
				}
			}
		}
	}
	return strconv.Itoa(count)
}

func part2() string {
	lines, err := util.ReadLines("data/day04/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	var rolls [][]bool

	for _, line := range lines {
		var row []bool
		for _, c := range line {
			row = append(row, c == '@')
		}
		rolls = append(rolls, row)
	}

	var directions [8][2]int = [8][2]int{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}

	count := 0
	moved := true
	for moved {
		count, moved = moveRolls(rolls, directions, count)
	}
	return strconv.Itoa(count)
}

func moveRolls(rolls [][]bool, directions [8][2]int, count int) (int, bool) {
	moved := false
	for y, row := range rolls {
		for x, val := range row {
			if val {
				neighbors := 0
				for _, direction := range directions {
					if check_cell(x+direction[0], y+direction[1], rolls) {
						neighbors += 1
					}
				}
				if neighbors < 4 {
					moved = true
					count += 1
					rolls[y][x] = false
				}
			}
		}
	}
	return count, moved
}

// slices are already built on top of a reference, so no pointers needed
func check_cell(x int, y int, grid [][]bool) bool {
	if x < 0 || x >= len(grid) ||
		y < 0 || y >= len(grid[0]) {
		return false
	}
	return grid[y][x]
}

func main() {
	answer1 := part1()
	fmt.Println("part 1:", answer1)
	answer2 := part2()
	fmt.Println("part 2:", answer2)
}
