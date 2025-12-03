package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("puzzle1_2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func part1(lines []string) int {
	dialStart := 50
	countZero := 0
	for _, line := range lines {
		n, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			dialStart = (dialStart + n) % 100
		} else {
			dialStart = (dialStart - n) % 100
			if dialStart < 0 {
				dialStart += 100
			}
		}
		if dialStart == 0 {
			countZero++
		}
	}
	return countZero
}

func part2(lines []string) int {
	dialStart := 50
	countZero := 0
	for _, line := range lines {
		n, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			for range n {
				dialStart--
				if dialStart == -1 {
					dialStart += 100
				}
				if dialStart == 0 {
					countZero++
				}
			}
		} else {
			for range n {
				dialStart++
				if dialStart == 100 {
					dialStart -= 100
				}
				if dialStart == 0 {
					countZero++
				}
			}
		}
	}
	return countZero
}
