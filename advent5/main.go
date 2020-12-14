package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	max := 0
	var wat = make([]bool, 955)
	for scanner.Scan() {
		var line = scanner.Text()
		num := findNum(line)
		if num > max {
			max = num
		}
		wat[num] = true
	}
	fmt.Println(max)
	for idx, exist := range wat {
		if !exist {
			fmt.Println(idx)
		}
	}
}

func findNum(line string) int {
	xLow := 0
	xHigh := 127
	yLow := 0
	yHigh := 7
	for _, char := range line {
		if char == 'F' {
			newX := (xLow + xHigh) / 2
			xHigh = newX
		} else if char == 'B' {
			newX := (xLow + xHigh) / 2
			xLow = newX + 1
		} else if char == 'R' {
			newY := (yLow + yHigh) / 2
			yLow = newY + 1
		} else if char == 'L' {
			newY := (yLow + yHigh) / 2
			yHigh = newY
		}
		// fmt.Println(xLow, xHigh, yLow, yHigh)
	}

	return 8*xLow + yLow
}
