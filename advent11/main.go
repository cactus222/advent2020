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

	seatMap := []string{}
	for scanner.Scan() {
		var line = scanner.Text()
		seatMap = append(seatMap, line)
	}
	var changed = true

	lenY := len(seatMap)
	lenX := len(seatMap[0])
	iter := 0
	for changed {
		changed, seatMap = transition(seatMap)
		iter++
		// if iter == 5 {
		// 	break
		// }
	}
	fmt.Println(iter)

	count := 0
	for i := 0; i < lenY; i++ {
		fmt.Println(seatMap[i])
		for j := 0; j < lenX; j++ {
			if seatMap[i][j] == '#' {
				count++
			}
		}
	}
	fmt.Println(count)
}

func transition(seatMap []string) (bool, []string) {

	lenY := len(seatMap)
	lenX := len(seatMap[0])
	var changed = false
	fmt.Println(lenY, lenX)
	var output = make([]string, 0)
	for y := 0; y < lenY; y++ {
		outputRow := ""
		for x := 0; x < lenX; x++ {

			if seatMap[y][x] == '.' {
				outputRow += "."
				continue
			} else if seatMap[y][x] == '#' {
				neighboursFilled := checkNeighbourCountsv2(seatMap, x, y)

				// neighboursFilled := checkNeighbourCounts(seatMap, x, y)
				// if neighboursFilled >= 4 {
				// 	outputRow += "L"
				// 	changed = true
				// } else {
				// 	outputRow += "#"
				// }
				if neighboursFilled >= 5 {
					outputRow += "L"
					changed = true
				} else {
					outputRow += "#"
				}

			} else if seatMap[y][x] == 'L' {
				neighboursFilled := checkNeighbourCountsv2(seatMap, x, y)
				//neighboursFilled := checkNeighbourCounts(seatMap, x, y)
				if neighboursFilled == 0 {
					outputRow += "#"
					changed = true
				} else {
					outputRow += "L"
				}
			}

		}
		output = append(output, outputRow)
	}
	return changed, output
}

func checkNeighbourCounts(seatMap []string, x int, y int) int {
	count := 0
	lenY := len(seatMap)
	lenX := len(seatMap[0])
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if y+i >= 0 && y+i < lenY && x+j >= 0 && x+j < lenX {
				if seatMap[y+i][x+j] == '#' {
					count++
				}
			}

		}
	}
	return count
}

func checkNeighbourCountsv2(seatMap []string, x int, y int) int {
	count := 0
	lenY := len(seatMap)
	lenX := len(seatMap[0])
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			for mult := 1; ; mult++ {
				deltai := i * mult
				deltaj := j * mult
				if y+deltai >= 0 && y+deltai < lenY && x+deltaj >= 0 && x+deltaj < lenX {
					if seatMap[y+deltai][x+deltaj] == '#' {
						count++
						break
					} else if seatMap[y+deltai][x+deltaj] == 'L' {
						break
					}
				} else {
					break
				}
			}

		}
	}
	return count
}
