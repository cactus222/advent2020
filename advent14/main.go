package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	var maskRegex = regexp.MustCompile(`mask = (\w+)`)
	var setRegex = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

	//index -> bit value
	var mask = make(map[int]bool, 0)
	var maskString = ""
	// idx -> value
	var memory = make(map[int64]int64, 0)
	for scanner.Scan() {
		var line = scanner.Text()

		var maskMatches = maskRegex.FindStringSubmatch(line)
		if len(maskMatches) > 0 {
			maskString = maskMatches[1]
			mask = make(map[int]bool, 0)
			for idx, cha := range maskString {
				realIdx := len(maskString) - 1 - idx
				if cha == '1' {
					mask[realIdx] = true
				} else if cha == '0' {
					mask[realIdx] = false
				}
			}
			// fmt.Println("maskva", maskString)
		} else {
			var setOpMatches = setRegex.FindStringSubmatch(line)
			var idx, _ = strconv.ParseInt(setOpMatches[1], 10, 64)
			var value, _ = strconv.ParseInt(setOpMatches[2], 10, 64)

			//part1
			// value = applyMask(mask, value)
			// memory[idx] = value
			indexes := applyMaskv2(maskString, idx)
			for _, idx := range indexes {
				memory[idx] = value
			}
		}

	}
	var sum int64 = 0
	for _, val := range memory {
		// fmt.Println(idx, val)
		sum += val
	}
	fmt.Println(sum)
}

func applyMaskv2(maskString string, value int64) []int64 {

	inputStr := pad(strconv.FormatInt(value, 2))
	outputStr := []byte{}
	for idx, cha := range maskString {
		if cha == '1' {
			outputStr = append(outputStr, '1')
		} else if cha == 'X' {
			outputStr = append(outputStr, 'X')
		} else {
			outputStr = append(outputStr, inputStr[idx])
		}
	}
	outputs := fillComboes(outputStr)

	return outputs
}

func fillComboes(outputStr []byte) []int64 {

	var outputs = []int64{}
	hasX := false
	for idx, cha := range outputStr {
		if cha == 'X' {
			outputStr[idx] = '0'
			outputs = append(outputs, fillComboes(outputStr)...)
			outputStr[idx] = '1'
			outputs = append(outputs, fillComboes(outputStr)...)
			outputStr[idx] = 'X'
			hasX = true
			break
		}
	}
	if !hasX {
		outputNum, _ := strconv.ParseInt(string(outputStr), 2, 64)
		return []int64{outputNum}
	}
	return outputs
}

func applyMask(mask map[int]bool, value int64) int64 {
	// fmt.Println(mask)
	// fmt.Println(" input", pad(strconv.FormatInt(value, 2)))
	for idx, isSet := range mask {
		// set idx
		if isSet {
			value |= 1 << idx
			// clear idx
		} else {
			value &= ^(1 << idx)
		}
	}
	// fmt.Println("output", pad(strconv.FormatInt(value, 2)))
	return value
}

func pad(in string) string {
	diff := 36 - len(in)
	return strings.Repeat("0", diff) + in
}
