// Go-based generator for creating new day scaffolding.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <day-number>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	arg := flag.Arg(0)
	// validate numeric
	if _, err := strconv.Atoi(arg); err != nil {
		fmt.Fprintf(os.Stderr, "day must be numeric\n")
		os.Exit(1)
	}

	n, _ := strconv.Atoi(arg)
	day := fmt.Sprintf("%02d", n)

	cmdDir := filepath.Join("cmd", "day"+day)
	dataDir := filepath.Join("data", "day"+day)

	// ensure not existing
	if exists(cmdDir) || exists(dataDir) {
		fmt.Fprintf(os.Stderr, "target already exists: %s or %s\n", cmdDir, dataDir)
		os.Exit(1)
	}

	// create directories
	if err := os.MkdirAll(cmdDir, 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "mkdir cmd: %v\n", err)
		os.Exit(1)
	}
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "mkdir data: %v\n", err)
		os.Exit(1)
	}

	// main.go template
	mainGo := `package main

import (
    "fmt"
    "log"

    "github.com/eliucid8/advent-of-code-2025/internal/util"
)

func part1(lines []string) string {
	return ""
}

func part2(lines []string) string {
	return ""
}

func main() {
    lines, err := util.ReadLines("data/day__PAD__/input.txt")
    if err != nil {
        log.Fatalf("read input: %v", err)
    }

	answer1 := part1(lines)
	fmt.Println("part 1:", answer1)
	answer2 := part2(lines)
	fmt.Println("part 2:", answer2)
}
`

	mainGo = strings.ReplaceAll(mainGo, "__PAD__", day)

	if err := os.WriteFile(filepath.Join(cmdDir, "main.go"), []byte(mainGo), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "write main.go: %v\n", err)
		os.Exit(1)
	}

	// create empty input files
	if err := os.WriteFile(filepath.Join(dataDir, "input.txt"), []byte(""), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "write input: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(filepath.Join(dataDir, "test-input.txt"), []byte(""), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "write test input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created %s and %s\n", cmdDir, dataDir)
}

func exists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
