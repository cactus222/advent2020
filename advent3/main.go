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

	// sizeLeft := 1
	var lines = []string{}
	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, line)
	}

	count1 := countLines(lines, 1, 1)
	count2 := countLines(lines, 3, 1)
	count3 := countLines(lines, 5, 1)
	count4 := countLines(lines, 7, 1)

	count5 := countLines(lines, 1, 2)
	fmt.Println(count1 * count2 * count3 * count4 * count5)

}

func countLines(lines []string, deltaX int, deltaY int) int {
	xIndex := 0
	i := 0
	size := 0
	count := 0
	for yIndex, line := range lines {

		if i == 0 || yIndex%deltaY != 0 {
			i += 1
			size = len(line)
			continue
		}
		xIndex += deltaX
		xIndex %= size
		if line[xIndex] == '#' {
			count++
		}
	}
	return count
}
