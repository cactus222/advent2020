package main

import (
	"bufio"
	"fmt"
	"os"
)

const NONE = 0
const MULT = 1
const ADD = 2

type Stuff struct {
	val int
	op  int
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		var line = scanner.Text()

		stack := [][]Stuff{}
		curStack := []Stuff{}
		for _, cha := range line {
			if cha == ' ' {
				// do nothing
			} else if cha == '(' {
				stack = append(stack, curStack)

				var newTempStack = make([]Stuff, 0)
				curStack = newTempStack
				// fmt.Println(stack)
				// fmt.Println("STACKS LEFT ", len(stack))

			} else if cha == ')' {
				evalOutput := eval(curStack)

				// fmt.Println("STACKS LEFT ", len(stack))
				curStack = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				curStack = append(curStack, evalOutput)
				// fmt.Println("STACK NOW", curStack)

			} else if cha == '*' {
				curStack = append(curStack, Stuff{0, MULT})
			} else if cha == '+' {
				curStack = append(curStack, Stuff{0, ADD})
			} else {
				value := int(cha - '0')
				//its a number
				curStack = append(curStack, Stuff{value, NONE})
			}
			// fmt.Println("SAW ", string(cha), "STACK NOW", curStack)
		}
		finalValue := eval(curStack)

		// fmt.Println(finalValue.val)
		total += finalValue.val
	}
	fmt.Println(total)
}

func eval(curStack []Stuff) Stuff {

	stack := []Stuff{}
	for idx, stuff := range curStack {
		// fmt.Println(stack)
		if idx == 0 {
			stack = append(stack, stuff)
			continue
		}
		//op
		if stuff.op != NONE {
			stack = append(stack, stuff)
		} else {
			//number
			rhs := stuff.val

			lastOp := NONE
			if len(stack) >= 1 {
				lastOp = stack[len(stack)-1].op
			}

			if lastOp == MULT {
				stack = append(stack, stuff)
				// just do it pt1

				// lastNum := stack[len(stack)-2].val
				// //pop last 2 values
				// stack = stack[:len(stack)-2]
				// //add back in addition
				// stack = append(stack, Stuff{lastNum + rhs, NONE})
			} else if lastOp == ADD {
				// fmt.Println(stack)
				lastNum := stack[len(stack)-2].val
				//pop last 2 values
				stack = stack[:len(stack)-2]
				//add back in addition
				stack = append(stack, Stuff{lastNum + rhs, NONE})
			}
		}
	}

	output := Stuff{1, NONE}
	for _, stuff := range stack {

		//op
		if stuff.op != NONE {
			// stack = append(stack, stuff)
		} else {
			//number
			rhs := stuff.val

			output.val *= rhs
		}
	}
	return output
}
