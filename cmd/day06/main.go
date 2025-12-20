package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/eliucid8/advent-of-code-2025/internal/util"
)

func part1(lines []string) string {
	var problems [][]int
	for _, line := range lines[:len(lines)-1] {
		var strnums = strings.Fields(line)
		var nums []int
		for _, num := range strnums {
			var newnum, _ = strconv.Atoi(num)
			if newnum != 0 {
				nums = append(nums, newnum)
			}
		}
		// fmt.Println(nums)
		// fmt.Println(len(nums))
		problems = append(problems, nums)
	}
	var signs = strings.Fields(lines[len(lines)-1])

	var total uint64
	for i := 0; i < len(problems[0]); i++ {
		var subtotal uint64
		if signs[i] == "*" {
			subtotal = 1
			for _, row := range problems {
				subtotal *= uint64(row[i])
			}
		} else {
			subtotal = 0
			for _, row := range problems {
				subtotal += uint64(row[i])
			}
		}
		total += subtotal
	}

	// fmt.Println(signs)
	// fmt.Println(len(signs))
	return strconv.FormatUint(total, 10)
}

func part2(lines []string) string {
	var signs = lines[len(lines)-1]
	col := 0
	var total uint64 = 0
	for col < len(signs) {
		var subtotal uint64
		if signs[col] == ' ' {
			col++
			continue
		}
		var multiply bool
		if signs[col] == '*' {
			subtotal = 1
			multiply = true
		} else {
			subtotal = 0
			multiply = false
		}

		for {
			var cur_num int
			for _, row := range lines[:len(lines)-1] {
				if row[col] != ' ' {
					cur_num *= 10
					cur_num += int(row[col] - '0')
				}
			}
			// fmt.Printf("%d ", cur_num)

			if cur_num == 0 {
				total += subtotal
				col++
				break
			} else {
				if multiply {
					subtotal *= uint64(cur_num)
					col++
				} else {
					subtotal += uint64(cur_num)
					col++
				}
			}
			if col == len(lines[0]) {
				total += subtotal
				break
			}
		}
		// fmt.Println()

	}

	return strconv.FormatUint(total, 10)
}

func main() {
	lines, err := util.ReadLines("data/day06/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	answer1 := part1(lines)
	fmt.Println("part 1:", answer1)
	answer2 := part2(lines)
	fmt.Println("part 2:", answer2)
}
