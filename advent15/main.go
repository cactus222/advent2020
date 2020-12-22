package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	val int64
	idx int64
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// number,idx
	var lastSpoken = map[int]int{}
	for scanner.Scan() {
		var line = scanner.Text()
		initial := strings.Split(line, ",")
		turn := 1
		for _, num := range initial {
			val, _ := strconv.ParseInt(num, 10, 64)
			lastSpoken[int(val)] = turn
			turn++
		}

		lastVal := 9999999
		for ; turn <= 30000000; turn++ {

			// fmt.Println(lastSpoken)
			startVal := lastVal
			if lastTurn, exists := lastSpoken[lastVal]; exists {
				deltaTurn := (turn - 1) - lastTurn
				lastSpoken[startVal] = (turn - 1)

				lastVal = deltaTurn
			} else {
				lastSpoken[startVal] = (turn - 1)
				lastVal = 0
			}
			// fmt.Println(lastVal)
		}
		fmt.Println(lastVal)
		break

	}
}
