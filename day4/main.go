package main

import (
	"adventofcode2025/internal/parse"
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

// helpers
var (
	directions = [8][2]int{ // most beautiful thing of the day
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /* {0,0} */, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
)

func toGrid(lines []string) [][]byte {
	grid := make([][]byte, len(lines))
	for i := range lines {
		grid[i] = []byte(lines[i])
	}
	return grid
}

func countNeighbors(grid [][]byte, i, j int) int {
	rows := len(grid)
	cnt := 0
	for _, d := range directions {
		ni, nj := i+d[0], j+d[1]
		if ni >= 0 && ni < rows && nj >= 0 && nj < len(grid[ni]) {
			if grid[ni][nj] == '@' {
				cnt++
			}
		}
	}
	return cnt
}

func canBeRemoved(grid [][]byte, i, j int) bool {
	if grid[i][j] != '@' {
		return false
	}
	return countNeighbors(grid, i, j) < 4
}

/*
Example:

..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
*/

func main() {
	part := flag.String("part", "both", "Which part to run: 1, 2, or both")
	inputFile := flag.String("input", "input.txt", "What test file should be used")
	flag.Parse()
	inputPath := filepath.Join("day4", *inputFile)

	lines, err := parse.ParseLines(inputPath)
	if err != nil {
		log.Fatalf("failed to parse input file: %v", err)
	}

	switch *part {
	case "1":
		fmt.Println("Part 1:", part1(lines))
	case "2":
		fmt.Println("Part 2:", part2(lines))
	case "both":
		fmt.Println("Part 1:", part1(lines))
		fmt.Println("Part 2:", part2(lines))
	default:
		log.Fatalf("unknown part: %s (expected 1, 2, or both)", *part)
	}
}

func part1(lines []string) any {
	grid := toGrid(lines)

	rowsLen := len(lines)
	colsLen := len(lines[0])

	removed := 0

	for ix := 0; ix < rowsLen; ix++ {
		for iy := 0; iy < colsLen; iy++ {
			if canBeRemoved(grid, ix, iy) {
				removed++
			}
		}
	}
	return removed
}

func helperPart2(grid [][]byte) int {
	rowsLen := len(grid)
	colsLen := len(grid[0])

	toRemove := make([][2]int, 0)

	for ix := 0; ix < rowsLen; ix++ {
		for iy := 0; iy < colsLen; iy++ {
			if canBeRemoved(grid, ix, iy) {
				toRemove = append(toRemove, [2]int{ix, iy})
			}
		}
	}
	for _, rem := range toRemove {
		grid[rem[0]][rem[1]] = '.'
	}
	return len(toRemove)
}

func recursiveRemove(grid [][]byte) int {
	removed := helperPart2(grid)
	if removed == 0 {
		return 0
	}
	return removed + recursiveRemove(grid)
}

func part2(lines []string) any {
	return recursiveRemove(toGrid(lines))
}
