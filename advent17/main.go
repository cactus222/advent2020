package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cube struct {
	x int
	y int
	z int
}

type HyperCube struct {
	x int
	y int
	z int
	w int
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cycles := 6
	row := 0
	delta := 0

	//pt1

	// actives := map[Cube]bool{}

	// for scanner.Scan() {
	// 	var line = scanner.Text()
	// 	for column, val := range line {
	// 		if val == '#' {
	// 			actives[Cube{column, row, 0}] = true
	// 		}
	// 		delta = column
	// 	}
	// 	row++
	// }
	// for i := 1; i <= cycles; i++ {
	// 	actives = transition(actives, i, delta)
	// 	fmt.Println(len(actives))
	// }

	actives := map[HyperCube]bool{}

	for scanner.Scan() {
		var line = scanner.Text()
		for column, val := range line {
			if val == '#' {
				actives[HyperCube{column, row, 0, 0}] = true
			}
			delta = column
		}
		row++
	}
	for i := 1; i <= cycles; i++ {
		actives = transitionHyper(actives, i, delta)
		fmt.Println(len(actives))
	}
}

func transition(actives map[Cube]bool, cycleNum, delta int) map[Cube]bool {
	newActives := make(map[Cube]bool, 0)

	for i := -cycleNum; i <= cycleNum+delta; i++ {
		for j := -cycleNum; j <= cycleNum+delta; j++ {
			for k := -cycleNum; k <= cycleNum; k++ {
				currentPoint := Cube{j, i, k}

				counts := checkNeighbourCounts(actives, currentPoint)
				isActive := actives[currentPoint]
				if isActive && (counts == 2 || counts == 3) {
					newActives[currentPoint] = true
				} else if !isActive && counts == 3 {
					newActives[currentPoint] = true
				}
				// fmt.Println("CHECK", j, i, k, isActive, counts)
			}
		}
	}
	return newActives
}

func checkNeighbourCounts(actives map[Cube]bool, currentPoint Cube) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				if actives[Cube{currentPoint.x + j, currentPoint.y + i, currentPoint.z + k}] {
					count++
				}
			}
		}
	}
	return count
}

func transitionHyper(actives map[HyperCube]bool, cycleNum, delta int) map[HyperCube]bool {
	newActives := make(map[HyperCube]bool, 0)

	for i := -cycleNum; i <= cycleNum+delta; i++ {
		for j := -cycleNum; j <= cycleNum+delta; j++ {
			for k := -cycleNum; k <= cycleNum; k++ {
				for l := -cycleNum; l <= cycleNum; l++ {
					currentPoint := HyperCube{j, i, k, l}

					counts := checkNeighbourCountsHyper(actives, currentPoint)
					isActive := actives[currentPoint]
					if isActive && (counts == 2 || counts == 3) {
						newActives[currentPoint] = true
					} else if !isActive && counts == 3 {
						newActives[currentPoint] = true
					}
				}
			}
		}
	}
	return newActives
}

func checkNeighbourCountsHyper(actives map[HyperCube]bool, currentPoint HyperCube) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue
					}
					if actives[HyperCube{currentPoint.x + j, currentPoint.y + i, currentPoint.z + k, currentPoint.w + l}] {
						count++
					}
				}
			}
		}
	}
	return count
}
