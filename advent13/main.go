package main

import (
	"bufio"
	"fmt"
	"os"
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

	idx := 0
	// var startTime int64 = 0
	//pt1

	// for scanner.Scan() {
	// 	var line = scanner.Text()
	// 	if idx == 0 {
	// 		startTime, _ = strconv.ParseInt(line, 10, 64)
	// 		idx++
	// 		continue
	// 	}

	// 	busids := strings.Split(line, ",")
	// 	var bestDelta int64 = 9999999
	// 	var bestID int64 = 0
	// 	for _, val := range busids {
	// 		value, err := strconv.ParseInt(val, 10, 64)
	// 		if err == nil {
	// 			delta := value - (startTime % value)
	// 			values = append(values, value)
	// 			fmt.Println(delta, value)
	// 			if delta < bestDelta {
	// 				bestDelta = delta
	// 				bestID = value
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(bestDelta, bestID)
	// 	fmt.Println(bestDelta * bestID)
	// }

	//pt2

	values := make([]pair, 0)
	for scanner.Scan() {
		var line = scanner.Text()
		if idx == 0 {
			// startTime, _ = strconv.ParseInt(line, 10, 64)
			idx++
			continue
		}

		busids := strings.Split(line, ",")

		for index, val := range busids {
			value, err := strconv.ParseInt(val, 10, 64)
			if err == nil {
				values = append(values, pair{
					val: value,
					idx: int64(index),
				})
			}
		}
	}
	fmt.Println(values)

	// values = []pair{pair{11, 6}, pair{16, 13}, pair{21, 9}, pair{25, 19}}
	//multiply all bus
	var product int64 = 1
	// assume all prime :/ else calc lcm
	for _, pair := range values {
		product *= pair.val
	}

	fmt.Println(product)

	// chinese remainder theorem dot dek
	var total int64 = 0
	for _, pair := range values {
		productWithoutSelf := product / pair.val
		modded := productWithoutSelf % pair.val

		multInverse := getMultInverse(modded, pair.val)
		multProductWithoutSelf := multInverse * productWithoutSelf
		cont := pair.idx * multProductWithoutSelf
		total += cont
		fmt.Println(productWithoutSelf, modded, multProductWithoutSelf, cont)
	}
	fmt.Println(total, product, total%product)
	fmt.Println(product - total%product)
}

// (modded * x) % pair.val = 1
func getMultInverse(modded int64, modolus int64) int64 {
	var guess int64
	for ; ; guess++ {
		if (modded*guess)%modolus == 1 {
			return guess
		}
	}
	return 0
}
