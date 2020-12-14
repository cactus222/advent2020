package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
)

func main() {


	var seenNums = make(map[int64]struct{});
	var seenNumsPairs = make(map[int64][]int64)
	
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		
		var num, _ = strconv.ParseInt(line, 10, 64)
		var remainder = 2020 - num

		if nums, exists := seenNumsPairs[remainder]; exists {
			fmt.Println(num * nums[0] * nums[1])
		} else {
			for seen, _ := range seenNums {
				seenNumsPairs[seen + num] = []int64{num, seen}
			}
			seenNums[num] = struct{}{}
		}
	}

}