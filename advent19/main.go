package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var rulesMap = make(map[int64][][]int64, 0)

	// 106: "a"
	// 65: "b"
	var stringsPart = false
	var targetStringsInputs = make([]string, 0)
	var maxTargetLength = 0
	//parse
	for scanner.Scan() {
		var line = scanner.Text()
		if stringsPart {
			targetStringsInputs = append(targetStringsInputs, line)
			if len(line) > maxTargetLength {
				maxTargetLength = len(line)
			}
		}

		ruleNum := strings.Split(line, ":")

		rule, err := strconv.ParseInt(ruleNum[0], 10, 64)
		if err != nil {
			stringsPart = true
			continue
		}

		rest := ruleNum[1]
		ruleMappingsStrs := strings.Split(rest, "|")
		rules := [][]int64{}
		for _, mappingsStr := range ruleMappingsStrs {
			numberStrings := strings.Split(strings.TrimSpace(mappingsStr), " ")
			nums := []int64{}
			for _, numString := range numberStrings {
				num, _ := strconv.ParseInt(numString, 10, 64)
				nums = append(nums, num)
			}
			rules = append(rules, nums)

		}
		rulesMap[rule] = rules
	}
	// fmt.Println(rulesMap)
	var rulesStringsMap = make(map[int64][]string, 0)
	rulesStringsMap[106] = []string{"a"}
	rulesStringsMap[65] = []string{"b"}
	// rulesStringsMap[4] = []string{"a"}
	// rulesStringsMap[5] = []string{"b"}
	// rulesStringsMap[1] = []string{"a"}
	// rulesStringsMap[14] = []string{"b"}
	redo := true

	// times8 := 0
	// times11 := 0
	// maxTimes := 4
	for redo {
		redo = false
		for rule, rules := range rulesMap {
			if rule == 0 {
				continue
			}

			canBeSolved := true
			// skip if we've already solved this unless its 8/11 OR 0
			if _, exists := rulesStringsMap[rule]; exists {
				// if _, exists := rulesStringsMap[rule]; exists {
				continue
			}

		OUTLOOP:
			for _, ruleNums := range rules {
				for _, ruleNum := range ruleNums {
					// ignore recursion non-solved
					if _, exists := rulesStringsMap[ruleNum]; !exists && ruleNum != rule {
						canBeSolved = false
						break OUTLOOP
					}
				}
			}
			if canBeSolved {
				// if rule == 8 {
				// 	times8++
				// 	if times8 >= maxTimes {
				// 		continue
				// 	}
				// }
				// if rule == 11 {
				// 	times11++
				// 	if times11 >= maxTimes {
				// 		continue
				// 	}
				// }

				fmt.Println(rule, rules)
				finalOutputs := crossJoin(rulesStringsMap, maxTargetLength, rules)

				if len(finalOutputs) != len(rulesStringsMap[rule]) {
					rulesStringsMap[rule] = finalOutputs
					//we updated something, go again
					redo = true
				}
			}
		}
	}

	fmt.Println(rulesStringsMap[31])
	fmt.Println(rulesStringsMap[42])
	r42Joined := strings.Join(rulesStringsMap[42], "|")
	r31Joined := strings.Join(rulesStringsMap[31], "|")

	fmt.Println(r42Joined)
	fmt.Println(r31Joined)

	// rulesStringsMap[0] = crossJoin(rulesStringsMap, maxTargetLength, rulesMap[0])
	// var rule0Strings = make(map[string]bool, 0)
	// for _, target := range rulesStringsMap[0] {
	// 	rule0Strings[target] = true
	// 	// fmt.Println(target)
	// }

	maxCopies := 10
	count := 0
	otherCount := 0
	for _, str := range targetStringsInputs {
		// if rule0Strings[str] {
		// 	count++
		// }
		for i := 1; i <= maxCopies; i++ {
			targetRegex := fmt.Sprintf("^(%s)+(%s){%d}(%s){%d}$", r42Joined, r42Joined, i, r31Joined, i)
			// fmt.Println(targetRegex)
			var curRegex = regexp.MustCompile(targetRegex)
			var matches = curRegex.FindStringSubmatch(str)
			if len(matches) > 0 {
				otherCount++
				break
			}

		}

	}
	fmt.Println("MAX LENGTH", maxTargetLength)
	fmt.Println("MATCHES", count)
	fmt.Println("MATCHESv2", otherCount)
}

func crossJoin(rulesStringsMap map[int64][]string, maxTargetLength int, rules [][]int64) []string {

	finalOutputs := []string{}
	for _, ruleNums := range rules {
		outputs := []string{}
		for idx, ruleNum := range ruleNums {
			if idx == 0 {
				values, _ := rulesStringsMap[ruleNum]
				outputs = append(outputs, values...)
			} else {
				//cross join
				values, _ := rulesStringsMap[ruleNum]
				newOutputsSet := map[string]bool{}
				for _, valueString := range values {
					for _, output := range outputs {
						output += valueString
						if len(output) <= maxTargetLength {
							newOutputsSet[output] = true
						}
						// newOutputs = append(newOutputs, output)
					}
				}
				newOutputs := make([]string, 0)
				for key, _ := range newOutputsSet {
					newOutputs = append(newOutputs, key)
				}
				outputs = newOutputs
			}
		}
		finalOutputs = append(finalOutputs, outputs...)
	}
	return finalOutputs
}
