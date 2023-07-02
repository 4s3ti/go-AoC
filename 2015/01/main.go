package main

import (
	"fmt"
	"os"
)

func elevator(instructions string) (int, int) {
	
	basementAt := 0
	floor := 0
	isBasement := false

	for pos, instruction := range instructions {
		if instruction == '(' {
			floor++
		} else if instruction == ')' {
			floor--
		}

		// Part 2
		if floor == -1 && !isBasement {
			isBasement = true
			basementAt = pos + 1
		}

	}
	return basementAt, floor
}


func main() {
	instructions, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	basementAt, floor := elevator(string(instructions))
	fmt.Println("Basement level reached for the first time at instruction:", basementAt)
	fmt.Println("Final floor:", floor)
}
