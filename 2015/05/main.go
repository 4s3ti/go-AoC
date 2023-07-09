package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hasTreeVowels(str string) bool {
	vowels := "aeiou"
	vowelsCount := 0
	hasTreeVowels := false
	for _, c := range str {
		if strings.ContainsAny(string(c), vowels) {
			vowelsCount++
		}
	}

	if vowelsCount >= 3 {
		hasTreeVowels = true
	}

	return hasTreeVowels
}

func containsBadStrings(str string) bool {
	badSTrings := []string{"ab", "cd", "pq", "xy"}
	hasBadStrings := false
	for i := 0; i < len(badSTrings); i++ {
		if strings.Contains(str, badSTrings[i]) {
			hasBadStrings = true
		}
	}

	return hasBadStrings
}

func hasRepeatedLetter(str string, pos int) bool {
	hasRepeatedLetter := false
	for i := 0; i < len(str)-pos; i++ {
		if str[i] == str[i+pos] {
			hasRepeatedLetter = true
		}
	}
	return hasRepeatedLetter
}

func hasRepeatedPairs(str string) bool {
	hasRepeatedPairs := false
	for i := 0; i < len(str)-2; i++ {
		if strings.Count(str, str[i:i+2]) >= 2 {
			hasRepeatedPairs = true
		}
	}

	return hasRepeatedPairs
}

func data() []string {
	data := []string{}
	d, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer d.Close()
	scanner := bufio.NewScanner(d)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func part1() {
	strings := data()
	totalNiceStrings := 0
	niceStrings := []string{}

	for i := 0; i < len(strings); i++ {
		if hasTreeVowels(strings[i]) && hasRepeatedLetter(strings[i], 1) && !containsBadStrings(strings[i]) {
			totalNiceStrings++
			niceStrings = append(niceStrings, strings[i])
		}
	}
	fmt.Println("Total nice strings: ", totalNiceStrings)
}

func part2() {
	strings := data()
	totalNiceStrings := 0

	for i := 0; i < len(strings); i++ {
		if hasRepeatedPairs(strings[i]) && hasRepeatedLetter(strings[i], 2) {
			totalNiceStrings++
		}
	}
	fmt.Println("Part 2, Total nice strings: ", totalNiceStrings)
}

func main() {
	part1()
	part2()
}
