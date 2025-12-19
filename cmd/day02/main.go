package main

import (
    "fmt"
    "log"

    "github.com/eliucid8/advent-of-code-25/internal/util"
)

func main() {
    lines, err := util.ReadLines("data/day02/input.txt")
    if err != nil {
        log.Fatalf("read input: %v", err)
    }
    fmt.Println("Line count:", len(lines))
    _ = lines
}
