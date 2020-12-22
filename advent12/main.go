package main

import (
	"bufio"
	"fmt"
	"math"
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

	var locX int64 = 0
	var locY int64 = 0
	var face int64 = 0 // East = 0. South = 1, West = 2, North = 3

	//pt1
	// for scanner.Scan() {
	// 	var line = scanner.Text()
	// 	var regex = regexp.MustCompile(`(\w)(\d+)`)

	// 	var matches = regex.FindStringSubmatch(line)
	// 	var op = matches[1]
	// 	val, _ := strconv.ParseInt(matches[2], 10, 64)

	// 	switch op {
	// 	case "N":
	// 		locY += val
	// 	case "S":
	// 		locY -= val
	// 	case "E":
	// 		locX += val
	// 	case "W":
	// 		locX -= val

	// 	case "F":
	// 		switch face {
	// 		case 0:
	// 			locX += val
	// 		case 1:
	// 			locY -= val
	// 		case 2:
	// 			locX -= val
	// 		case 3:
	// 			locY += val
	// 		}
	// 	case "R":
	// 		face = (face + 4 + val/90) % 4
	// 	case "L":
	// 		face = (face + 4 - val/90) % 4
	// 	}

	// }
	// fmt.Println(locX, locY, face, math.Abs(float64(locX))+math.Abs(float64(locY)))

	//pt2

	var wayX int64 = 10
	var wayY int64 = 1
	for scanner.Scan() {
		var line = scanner.Text()
		var regex = regexp.MustCompile(`(\w)(\d+)`)

		var matches = regex.FindStringSubmatch(line)
		var op = matches[1]
		val, _ := strconv.ParseInt(matches[2], 10, 64)

		switch op {
		case "N":
			wayY += val
		case "S":
			wayY -= val
		case "E":
			wayX += val
		case "W":
			wayX -= val

		case "F":
			locX += val * wayX
			locY += val * wayY
		case "R":
			val = val / 90
			if val == 1 {
				temp := wayX
				wayX = wayY
				wayY = -temp
			} else if val == 2 {
				wayX *= -1
				wayY *= -1
			} else if val == 3 {
				temp := wayX
				wayX = -wayY
				wayY = temp
			} else {
				fmt.Println("FIXME")
			}
		case "L":
			val = val / 90
			if val == 1 {
				temp := wayX
				wayX = -wayY
				wayY = temp
			} else if val == 2 {
				wayX *= -1
				wayY *= -1
			} else if val == 3 {
				temp := wayX
				wayX = wayY
				wayY = -temp
			} else {
				fmt.Println("FIXME")
			}
		}

	}
	fmt.Println(locX, locY, face, math.Abs(float64(locX))+math.Abs(float64(locY)))

}
