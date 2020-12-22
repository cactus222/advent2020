package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numbers := []int64{0}
	for scanner.Scan() {
		var line = scanner.Text()
		val, _ := strconv.ParseInt(line, 10, 64)

		numbers = append(numbers, val)
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	numbers = append(numbers, numbers[len(numbers)-1]+3)

	var deltas = []int{0, 0, 0, 0}
	for idx, v := range numbers {
		if idx > 0 {

			deltas[v-numbers[idx-1]]++
		}
	}
	fmt.Println(deltas[1], deltas[2], deltas[3], deltas[1]*deltas[3])

	//math is hard.
	var variations = 1
	var delta1Streak = 1
	for idx, v := range numbers {

		if idx == 0 {
			continue
		}
		var delta = v - numbers[idx-1]
		if delta == 3 {
			fmt.Println(delta1Streak)
			// multiply variations
			if delta1Streak == 3 {
				variations *= 2
			} else if delta1Streak == 4 {
				variations *= 4
			} else if delta1Streak == 5 {
				variations *= 7
			}
			delta1Streak = 1
		} else if delta == 1 {
			delta1Streak++
		}
	}

	fmt.Println(variations)
}
