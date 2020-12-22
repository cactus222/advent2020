package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Interval struct {
	low  int64
	high int64
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var maskRegex = regexp.MustCompile(`(.*): (\d+)-(\d+) or (\d+)-(\d+)`)

	var mappings = make(map[string][]*Interval, 0)

	var sum int64 = 0
	var myTicket = []int64{}
	var validTickets = [][]int64{}
	for scanner.Scan() {
		var line = scanner.Text()
		var maskMatches = maskRegex.FindStringSubmatch(line)
		if len(maskMatches) > 0 {
			low1, _ := strconv.ParseInt(maskMatches[2], 10, 64)
			high1, _ := strconv.ParseInt(maskMatches[3], 10, 64)
			low2, _ := strconv.ParseInt(maskMatches[4], 10, 64)
			high2, _ := strconv.ParseInt(maskMatches[5], 10, 64)
			intervals := []*Interval{&Interval{low1, high1}, &Interval{low2, high2}}
			mappings[maskMatches[1]] = intervals
		} else {
			numbers := strings.Split(line, ",")
			if len(numbers) > 1 {
				ticket := []int64{}
				valid := true
				for _, numString := range numbers {
					num, _ := strconv.ParseInt(numString, 10, 64)
					ticket = append(ticket, num)
					if !isValid(num, mappings) {
						sum += num
						valid = false
					}
				}
				if valid {
					validTickets = append(validTickets, ticket)
				}
				if len(myTicket) == 0 {
					myTicket = ticket
				}
			}
		}
	}
	fmt.Println(myTicket)
	// field -> idx -> possible
	//everything starts possible
	var possibilities = make(map[string][]bool, 0)
	for key, _ := range mappings {
		var startMap = make([]bool, len(mappings))
		for i := 0; i < len(mappings); i++ {
			startMap[i] = true
		}
		possibilities[key] = startMap
	}

	// check for every ticket, if that idx is valid for that slot
	for key, _ := range mappings {
		possibilityArray := possibilities[key]

		for _, ticket := range validTickets {
			for idx, value := range ticket {
				if possibilityArray[idx] && !isValidForKey(value, key, mappings) {
					possibilityArray[idx] = false
				}
			}
		}
	}

	fmt.Println("BEFORE MOD")
	for key, possibleIndexes := range possibilities {
		fmt.Println(key, possibleIndexes)
	}

	fmt.Println("MODDING")
	redo := true
	for redo {
		redo = false
		for _, possibleIndexes := range possibilities {
			numTrues := 0
			theTrueIndex := 0
			for idx, possible := range possibleIndexes {
				if possible {
					numTrues++
					theTrueIndex = idx
				}
			}
			// this is the only valid one
			if numTrues == 1 {
				// modify all others to false
				for _, modPossibleIndexes := range possibilities {
					modPossibleIndexes[theTrueIndex] = false
				}
			} else if numTrues > 1 {
				redo = true
			}

			possibleIndexes[theTrueIndex] = true
		}
	}

	var finalMapping = make([]string, len(mappings))

	fmt.Println("AFTER MOD")
	for key, possibleIndexes := range possibilities {
		fmt.Println(key, possibleIndexes)
		for idx, possible := range possibleIndexes {
			if possible {
				finalMapping[idx] = key
			}
		}
	}

	var product int64 = 1
	fmt.Println("FINAL MAPPING")
	for idx, key := range finalMapping {
		fmt.Println(idx, key, myTicket[idx])
		if strings.HasPrefix(key, "departure") {
			product *= myTicket[idx]
		}
	}
	fmt.Println("FINAL PRODUCT", product)
	//pt1
	// for key, intervals := range mappings {
	// 	for _, interval := range intervals {
	// 		fmt.Println(key, interval)
	// 	}
	// }
	// fmt.Println(sum)
}

func isValid(num int64, mappings map[string][]*Interval) bool {
	for _, intervals := range mappings {
		for _, interval := range intervals {
			if num >= interval.low && num <= interval.high {
				return true
			}
		}
	}
	return false
}

func isValidForKey(num int64, key string, mappings map[string][]*Interval) bool {
	intervals := mappings[key]

	for _, interval := range intervals {
		if num >= interval.low && num <= interval.high {
			return true
		}
	}

	return false
}
