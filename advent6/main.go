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

	chars := make(map[rune]int, 0)
	count := 0
	people := 0
	for scanner.Scan() {
		var line = scanner.Text()

		if len(line) == 0 {
			for _, val := range chars {
				if val == people {
					count++
				}
			}
			// count += len(chars)
			people = 0
			chars = make(map[rune]int, 0)
		} else {
			for _, x := range line {
				chars[x] = chars[x] + 1
			}
			people++
		}

	}
	fmt.Println(count)

}
