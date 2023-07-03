package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"
	"sort"
)

type Box struct {
	Lenght int
	Width int
	Height int
}

func (b Box) NeededPaper() int {
	sidesArea := []int{2 * b.Lenght * b.Width, 2 * b.Width * b.Height, 2 * b.Height * b.Lenght}
	sort.Ints(sidesArea)
	slack := sidesArea[0] / 2
	return  sidesArea[0] + sidesArea[1] + sidesArea[2] + slack
}

func (b Box) NeededRibbon() int {
	measures := []int{b.Lenght, b.Width, b.Height}
	sort.Ints(measures)
	return 2 * measures[0] + 2 * measures[1] + b.Lenght * b.Width * b.Height
}

func main() {
	totalPaper := 0
	totalRibbon := 0
	boxes, err := os.Open("input.txt")
	if err != nil {
	  panic(err)
	}
	defer boxes.Close()
	scanner := bufio.NewScanner(boxes)
	for scanner.Scan() {
		measures := strings.Split(scanner.Text(), "x")
		l, _ := strconv.Atoi(measures[0])
		w, _ := strconv.Atoi(measures[1])
		h, _ := strconv.Atoi(measures[2])
		box := Box{l, w, h}
		totalPaper += box.NeededPaper()
		totalRibbon += box.NeededRibbon()
	}
	fmt.Println("Total paper needed: ",  totalPaper)
	fmt.Println("Total riboon needed: ", totalRibbon)
}
