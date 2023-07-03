package main

import (
	"fmt"
	"os"
)

func elevator(instructions string) int{
	var floor int
	for _, instruction := range instructions {
		if instruction == '(' {
			floor++
		} else if instruction == ')' {
			floor--
		}
	}

	return  floor
}

// This approach returns the wrong answer
// func doesntWork() {
// 	var instructions string
// 	fmt.Print("What floor you want to go?: " )
// 	fmt.Scan(&instructions)
//
//
//
// 	lastFloor := elevator(instructions)
// 	fmt.Println("final Floor: ", lastFloor)
// }

// Let's find our why
func findWhy() {
	var instructionsFromInput string
	fmt.Print("What floor you want to go?: " )
	fmt.Scan(&instructionsFromInput)
	fmt.Println("instructions: ", instructionsFromInput)
	fmt.Println("Length from input: ", len(instructionsFromInput))

	instructionsFromFile, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Lenght from file: ", len(instructionsFromFile))
	lastFloorFromInput := elevator(instructionsFromInput)
	lastFloorFromFile := elevator(string(instructionsFromFile))
	fmt.Println("Final Floor From input: ", lastFloorFromInput)
	fmt.Println("Final Floor from file: ", lastFloorFromFile)
}

func main() {
	// doesntWork()
	findWhy()
}
