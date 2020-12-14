package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	valid := 0
	for scanner.Scan() {
		var line = scanner.Text()
		var regex = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

		// only expect first match
		var parsed = regex.FindAllStringSubmatch(line, -1)[0]

		// is 1 indexed
		min, _ := strconv.ParseInt(parsed[1], 10, 64)
		max, _ := strconv.ParseInt(parsed[2], 10, 64)
		targetChar := parsed[3][0]
		targetString := parsed[4]

		if isValidPart2(targetChar, targetString, min, max) {
			valid++
		}
	}
	fmt.Println(valid)

}

func isValidPart1(targetChar byte, targetString string, min int64, max int64) bool {
	var count int64 = 0
	for _, char := range targetString {
		if byte(char) == targetChar {
			count++
		}
	}

	return (count >= min) && count <= max
}

func isValidPart2(targetChar byte, targetString string, min int64, max int64) bool {
	first := byte(targetString[min-1]) == targetChar
	second := byte(targetString[max-1]) == targetChar

	return (!first && second) || (!second && first)
}
