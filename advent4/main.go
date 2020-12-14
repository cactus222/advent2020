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

	currentFields := make(map[string]string, 0)
	valid := 0
	var reqFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for scanner.Scan() {
		var line = scanner.Text()
		var regex = regexp.MustCompile(`(\w+):(\S+)`)

		// only expect first match
		var parsed = regex.FindAllStringSubmatch(line, -1)

		for _, groupMatch := range parsed {
			// for idx, field := range groupMatch {
			// 	if idx == 0 {
			// 		continue // skip main
			// 	}
			// 	if idx == 1 {
			// 		currentFields[field] = struct{}{}
			// 	}

			// 	fmt.Println(field)
			// }
			currentFields[groupMatch[1]] = groupMatch[2]
		}

		if len(parsed) == 0 {
			if isValid2(currentFields, reqFields) {
				valid++
				fmt.Println("VALID ", currentFields)
			} else {
				fmt.Println("INVALID ", currentFields)
			}
			currentFields = make(map[string]string, 0)
		}

	}
	fmt.Println(valid)

}

func isValid(currentFields map[string]string, reqFields []string) bool {
	for _, field := range reqFields {
		if _, exists := currentFields[field]; !exists {
			return false
		}
	}
	return true
}

func isValid2(currentFields map[string]string, reqFields []string) bool {
	var heightCMRegex = regexp.MustCompile(`(\d{3})cm`)
	var heightINRegex = regexp.MustCompile(`(\d{2})in`)
	var hairRegex = regexp.MustCompile(`#[0-9a-z]{6}`)
	var pidRegex = regexp.MustCompile(`\d{9}`)
	// too lazy to make fcns
	for _, field := range reqFields {
		if val, exists := currentFields[field]; !exists {
			return false
		} else {
			if field == "byr" {
				val, err := strconv.ParseInt(val, 10, 64)
				if err != nil {
					fmt.Println("byr err", val)
					return false
				}
				if !(val >= 1920 && val <= 2002) {
					fmt.Println("byr bad range", val)
					return false
				}
			} else if field == "iyr" {
				val, err := strconv.ParseInt(val, 10, 64)
				if err != nil {
					fmt.Println("iyr err", val)
					return false
				}
				if !(val >= 2010 && val <= 2020) {
					fmt.Println("iyr bad range", val)
					return false
				}
			} else if field == "eyr" {
				val, err := strconv.ParseInt(val, 10, 64)
				if err != nil {
					fmt.Println("eyr invalid", val)
					return false
				}
				if !(val >= 2020 && val <= 2030) {
					fmt.Println("eyr bad range", val)
					return false
				}
			} else if field == "hgt" {
				var parsed = heightCMRegex.FindStringSubmatch(val)
				heightType := "cm"
				if len(parsed) == 0 || len(val) != 5 {
					parsed = heightINRegex.FindStringSubmatch(val)
					heightType = "in"
				}
				if len(parsed) > 0 && ((heightType == "in" && len(val) == 4) || (heightType == "cm" && len(val) == 5)) {
					num := parsed[1]
					heightValue, err := strconv.ParseInt(num, 10, 64)
					if err != nil {
						fmt.Println("Hgt err", val)
						return false
					}

					if heightType == "cm" {
						if !(heightValue >= 150 && heightValue <= 193) {
							fmt.Println("Hgt cm range", val)
							return false
						}
					} else if heightType == "in" {
						if !(heightValue >= 59 && heightValue <= 76) {
							fmt.Println("Hgt in range", val)
							return false
						}
					} else {
						return false
					}
				} else {
					return false
				}
			} else if field == "hcl" {
				var parsed = hairRegex.FindStringSubmatch(val)
				if len(parsed) == 0 || len(val) != 7 {
					fmt.Println("hcl err", val)
					return false
				}
			} else if field == "ecl" {
				validEyes := map[string]int{"amb": 1, "blu": 1, "brn": 1, "gry": 1, "grn": 1, "hzl": 1, "oth": 1}
				if validEyes[val] == 0 {
					fmt.Println("ecl err", val)
					return false
				}
			} else if field == "pid" {
				var parsed = pidRegex.FindStringSubmatch(val)
				if len(parsed) == 0 || len(val) != 9 {

					fmt.Println("pid err", val)
					return false
				}
			}
		}
	}
	return true
}
