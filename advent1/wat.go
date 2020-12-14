package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
)

func main() {


	var seenNums = make(map[int64]struct{});
	
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		
		var num, _ = strconv.ParseInt(line, 10, 64)
		if _, exists := seenNums[2020 - num]; exists {

			fmt.Println((2020 - num) * num)
		} else {
			seenNums[num] = struct{}{}
		}
	}

}