package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		lines = strings.Split(line, ",")
	}

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))

}

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		numbersRange := strings.Split(line, "-")

		if len(numbersRange) != 2 {
			panic("invalid line") // this shouldn't happen
		}
		leftLimit, _ := strconv.Atoi(numbersRange[0])
		rightLimit, _ := strconv.Atoi(numbersRange[1])
		for currNum := leftLimit; currNum <= rightLimit; currNum++ {
			currNumString := strconv.Itoa(currNum)
			if lenCurrString := len(currNumString); lenCurrString%2 == 0 && currNumString[:lenCurrString/2] == currNumString[lenCurrString/2:] {
				total += currNum
			}
		}
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		numbersRange := strings.Split(line, "-")
		if len(numbersRange) != 2 {
			panic("invalid line")
		}
		leftLimit, _ := strconv.Atoi(numbersRange[0])
		rightLimit, _ := strconv.Atoi(numbersRange[1])
		for currNum := leftLimit; currNum <= rightLimit; currNum++ {
			currNumString := strconv.Itoa(currNum)
			for size := 1; size <= len(currNumString)/2; size++ {
				if len(currNumString)%size != 0 {
					continue
				}
				pattern := currNumString[:size]
				repeated := true
				for ix := size; ix < len(currNumString); ix += size {
					if ix+size > len(currNumString) {
						repeated = false
						break
					}
					if currNumString[ix:ix+size] != pattern {
						repeated = false
						break
					}
				}
				if repeated {
					total += currNum
					break
				}
			}
		}
	}
	return total
}
