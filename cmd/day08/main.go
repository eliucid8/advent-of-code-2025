package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/eliucid8/advent-of-code-2025/internal/util"
)

func part1(lines []string) string {
	const NUM_CONNECTIONS int = 1000
	boxes := parseBoxes(lines)

	shortest_edges := shortest_n_edges(boxes, NUM_CONNECTIONS)

	// fmt.Println(shortest_edges)

	var uf []int
	for i := range len(boxes) {
		uf = append(uf, i)
	}
	var tree_size []int
	for range len(boxes) {
		tree_size = append(tree_size, 1)
	}

	// perform union-finding
	for _, edge := range shortest_edges {
		uf_merge(edge[0], edge[1], uf[:], tree_size[:])
		// fmt.Println(tree_size)
	}
	sort.Ints(tree_size[:])
	slices.Reverse(tree_size)

	// fmt.Println(tree_size)
	return strconv.Itoa(tree_size[0] * tree_size[1] * tree_size[2])
}

func shortest_n_edges(boxes [][3]int, n int) [][2]int {
	var shortest_edges [][2]int
	var dists []float64
	for a := 0; a < len(boxes)-1; a++ {
		for b := a + 1; b < len(boxes); b++ {
			var cur_len = dist(boxes[a], boxes[b])
			if len(dists) == 0 {
				shortest_edges = append(shortest_edges, [2]int{a, b})
				dists = append(dists, cur_len)
				continue
			}

			if cur_len > dists[len(dists)-1] {
				if len(shortest_edges) < n {
					shortest_edges = append(shortest_edges, [2]int{a, b})
					dists = append(dists, cur_len)
				}
				continue
			} else {
				shortest_edges[len(shortest_edges)-1] = [2]int{a, b}
				dists[len(dists)-1] = cur_len
			}

			for i := len(dists) - 1; i >= 1 && cur_len < dists[i-1]; i-- {
				shortest_edges[i-1], shortest_edges[i] = shortest_edges[i], shortest_edges[i-1]
				dists[i-1], dists[i] = dists[i], dists[i-1]
			}
		}
	}
	return shortest_edges
}

func parseBoxes(lines []string) [][3]int {
	var boxes [][3]int
	for _, line := range lines {
		var scoords = strings.Split(line, ",")
		var icoords [3]int
		for i, scoord := range scoords {
			icoords[i], _ = strconv.Atoi(scoord)
		}
		boxes = append(boxes, icoords)
	}
	return boxes
}

func uf_find_parent(i int, uf []int) int {
	var intermediate []int
	intermediate = append(intermediate, i)
	for uf[i] != i {
		intermediate = append(intermediate, i)
		i = uf[i]
	}
	for _, node := range intermediate {
		uf[node] = i
	}

	return i
}

func uf_merge(a int, b int, uf []int, sizes []int) {
	a = uf_find_parent(a, uf)
	b = uf_find_parent(b, uf)
	if a == b {
		return
	}

	if sizes[a] >= sizes[b] {
		uf[b] = a
		sizes[a] += sizes[b]
		sizes[b] = 0
	} else {
		uf[a] = b
		sizes[b] += sizes[a]
		sizes[a] = 0
	}
}

func dist(a [3]int, b [3]int) float64 {
	var d0 = (a[0] - b[0]) * (a[0] - b[0])
	var d1 = (a[1] - b[1]) * (a[1] - b[1])
	var d2 = (a[2] - b[2]) * (a[2] - b[2])
	return math.Sqrt(float64(d0 + d1 + d2))
}

func part2(lines []string) string {
	// oops ig I gotta tweak this
	const NUM_CONNECTIONS int = 10000
	boxes := parseBoxes(lines)

	shortest_edges := shortest_n_edges(boxes, NUM_CONNECTIONS)

	// fmt.Println(shortest_edges)

	var uf []int
	for i := range len(boxes) {
		uf = append(uf, i)
	}
	var tree_size []int
	for range len(boxes) {
		tree_size = append(tree_size, 1)
	}

	// perform union-finding
	for _, edge := range shortest_edges {
		uf_merge(edge[0], edge[1], uf[:], tree_size[:])
		var a_size = tree_size[uf[edge[0]]]
		var b_size = tree_size[uf[edge[1]]]
		if max(a_size, b_size) == len(boxes) {
			var box_xs = boxes[edge[0]][0] * boxes[edge[1]][0]
			return strconv.Itoa(box_xs)
		}
		// fmt.Println(tree_size)
	}

	// fmt.Println(tree_size)
	return "Not enough connections. Please increase NUM_CONNECTIONS."
}

func main() {
	lines, err := util.ReadLines("data/day08/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	answer1 := part1(lines)
	fmt.Println("part 1:", answer1)
	answer2 := part2(lines)
	fmt.Println("part 2:", answer2)
}
