package main

import ( 
	"fmt"
	"strconv"
	"os"
	"bufio"
)


func instructions() []string {
	instructions := []string{}

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	return instructions
}

func Part1(instructions []string) int { 
	rawStringLen := 0
	unquotedStringlen := 0
	for i := 0; i < len(instructions); i++ {
		rawStringLen += len(instructions[i])
		unquoted, _ := strconv.Unquote(instructions[i])
		unquotedStringlen += len(unquoted)
	}

	return rawStringLen - unquotedStringlen
}

func Part2(instructions []string) int {
	encodedStringLen := 0
	rawStringLen := 0
	for i := 0; i < len(instructions); i++ {
		encodedStringLen += len(strconv.Quote(instructions[i]))
		rawStringLen += len(instructions[i])
	}

	return encodedStringLen - rawStringLen
}

func main() {
	instructions := instructions()
	fmt.Println("Part 1: ", Part1(instructions))
	fmt.Println("Part 2: ", Part2(instructions))
}

