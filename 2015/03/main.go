package main

import (
	"bufio"
	"fmt"
	"os"
)

func instructions() string {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var instructions string
	for scanner.Scan() {
		instructions += scanner.Text()
	}
	return instructions
}

func movement(instruction string) (int, int) {
	x := 0
	y := 0
	for _, movement := range instruction {
		switch movement {
		case '^': y++
		case 'v': y--
		case '>': x++
		case '<': x--
		default:
			fmt.Printf("%v is an invalid movement", string(movement))
		}
	}
	return x, y
}

func uniqueCoordinates(coordinates [][2]int) map[[2]int]bool {
	ucoords := make(map[[2]int]bool)
	for _, coordinate := range coordinates {
		if !ucoords[coordinate] {
			ucoords[coordinate] = true
		}
	}

	return ucoords
}

func part1(instructions string) {
	currentx := 0
	currenty := 0
	coordinates := [][2]int{{0, 0}}
	for _, instruction := range instructions {
		x, y := movement(string(instruction))
		currentx += x
		currenty += y
		coordinates = append(coordinates, [2]int{currentx, currenty})
	}

	uniqueDeliveries := uniqueCoordinates(coordinates)
	fmt.Println(len(uniqueDeliveries))
}

func part2(instructions string) {
	santaTurn := true
	currentSantaX := 0
	currentSantaY := 0
	currentRoboSantaX := 0
	currentRoboSantaY := 0

	coordinates := [][2]int{{0, 0}, {0, 0}}
	for _, instruction := range instructions {
		if santaTurn {
			x, y := movement(string(instruction))
			currentSantaX += x
			currentSantaY += y
			coordinates = append(coordinates, [2]int{currentSantaX, currentSantaY})
			santaTurn = false
		} else {
			x, y := movement(string(instruction))
			currentRoboSantaX += x
			currentRoboSantaY += y
			coordinates = append(coordinates, [2]int{currentRoboSantaX, currentRoboSantaY})
			santaTurn = true
		}
	}

	uniqueDeliveries := uniqueCoordinates(coordinates)
	fmt.Println(len(uniqueDeliveries))
}

func main() {
	instructions := instructions()
	part1(instructions)
	part2(instructions)
}
