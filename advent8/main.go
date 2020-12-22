package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Instruction struct {
	op      string
	val     int64
	visited bool
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instrs := []*Instruction{}

	for scanner.Scan() {
		var line = scanner.Text()
		var regex = regexp.MustCompile(`(.+) (.+)`)

		var matches = regex.FindStringSubmatch(line)
		val, _ := strconv.ParseInt(matches[2], 10, 64)
		op := matches[1]
		instr := &Instruction{
			op:      op,
			val:     val,
			visited: false,
		}
		instrs = append(instrs, instr)
	}

	clear := func() {
		for _, instr := range instrs {
			instr.visited = false
		}
	}
	for _, instr := range instrs {
		if instr.op == "jmp" {
			instr.op = "nop"
			if tryFcn(instrs) {
				break
			}
			instr.op = "jmp"
			clear()
		} else if instr.op == "nop" {
			instr.op = "jmp"
			if tryFcn(instrs) {
				break
			}
			instr.op = "nop"
			clear()
		}
	}

}

func tryFcn(instrs []*Instruction) bool {
	var pc int64 = 0
	var acc int64 = 0
	for {
		if pc >= int64(len(instrs)) {
			fmt.Println("SUCCESS")
			fmt.Println("acc", acc)
			return true

		}
		instr := instrs[pc]
		if instr.visited {
			fmt.Println("acc", acc)
			return false
		}
		instr.visited = true

		if instr.op == "jmp" {
			pc += instr.val
		} else if instr.op == "acc" {
			acc += instr.val
			pc++
		} else {
			pc++
		}
	}

}
