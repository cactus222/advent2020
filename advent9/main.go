package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Instruction struct {
	op      string
	val     int64
	visited bool
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numbers := []int64{}
	idx := 0
	var badVal int64 = 0
	for scanner.Scan() {
		var line = scanner.Text()
		val, _ := strconv.ParseInt(line, 10, 64)
		if idx >= 25 {
			if !checkLast25(numbers, val, idx) {
				fmt.Println(val, idx)
				badVal = val
				break
			}
		}
		numbers = append(numbers, val)
		idx++
	}
	idx--
	var total int64 = 0
	highIndex := idx

	for ; ; idx-- {
		total += numbers[idx]
		if total == badVal {
			fmt.Println("DONE, range is ", highIndex, idx)
			fmt.Println(numbers[idx : highIndex+1])
			//sum max+min
			break
		}
		if total > badVal {
			total -= numbers[highIndex]
			highIndex--
		}
	}

}

func checkLast25(numbers []int64, val int64, idx int) bool {
	chunk := numbers[idx-25 : idx]

	var seenNums = make(map[int64]struct{})
	for _, x := range chunk {

		if _, exists := seenNums[val-x]; exists {

			return true
		} else {
			seenNums[x] = struct{}{}
		}
	}
	return false
}
