package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Matrix struct {
	coords [1000][1000]int
}

func (m *Matrix) TurnOn(x int, y int) {
	m.coords[x][y] = 1
}

func (m *Matrix) TurnOff(x int, y int) {
	m.coords[x][y] = 0
}

func (m *Matrix) ShowMatrix() {
	for y := 0; y < len(m.coords); y++ {
		for x := 0; x < len(m.coords[y]); x++ {
			fmt.Print(m.coords[y][x], " ")
		}
		fmt.Println()
	}
}

func (m *Matrix) TurnOnGrid(from [2]int, to [2]int) {
	for y := from[1]; y <= to[1]; y++ {
		for x := from[0]; x <= to[0]; x++ {
			m.coords[x][y] += 1
		}
	}
}

func (m *Matrix) TurnOffGrid(from [2]int, to [2]int) {
	for y := from[1]; y <= to[1]; y++ {
		for x := from[0]; x <= to[0]; x++ {
			if m.coords[x][y] == 0 {
				continue
			} else {
				m.coords[x][y] -= 1
			}
		}
	}
}

func (m *Matrix) Toggle(from [2]int, to [2]int) {
	for y := from[1]; y <= to[1]; y++ {
		for x := from[0]; x <= to[0]; x++ {
			if m.coords[x][y] == 1 {
				m.TurnOff(x, y)
			} else {
				m.TurnOn(x, y)
			}
		}
	}
}

func (m *Matrix) LightCount() int {
	count := 0
	for y := 0; y < len(m.coords); y++ {
		for x := 0; x < len(m.coords[y]); x++ {
			if m.coords[x][y] != 0 {
				count++
			}
		}
	}
	return count
}

func (m *Matrix) TotalBrightness() int {
	count := 0
	for y := 0; y < len(m.coords); y++ {
		for x := 0; x < len(m.coords[y]); x++ {
			count += m.coords[x][y]
		}
	}
	return count

}

func instructions() []string {
	instructions := []string{}
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	return instructions
}

func parseCoords(s string) [2]int {
	coords := strings.Split(s, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return [2]int{x, y}
}

func part1() {
	m := Matrix{}

	instructions := instructions()
	for i := 0; i < len(instructions); i++ {
		inst := strings.Split(instructions[i], " ")
		if inst[0] == "turn" {
			from := parseCoords(inst[2])
			to := parseCoords(inst[4])
			if inst[1] == "on" {
				m.TurnOnGrid(from, to)
			}

			if inst[1] == "off" {
				m.TurnOffGrid(from, to)
			}
		}

		if inst[0] == "toggle" {
			from := parseCoords(inst[1])
			to := parseCoords(inst[3])
			m.Toggle(from, to)
		}
	}
	fmt.Println(m.LightCount())
}

func part2() {
	m := Matrix{}

	instructions := instructions()
	for i := 0; i < len(instructions); i++ {
		inst := strings.Split(instructions[i], " ")
		if inst[0] == "turn" {
			from := parseCoords(inst[2])
			to := parseCoords(inst[4])
			if inst[1] == "on" {
				m.TurnOnGrid(from, to)
			}

			if inst[1] == "off" {
				m.TurnOffGrid(from, to)
			}
		}

		if inst[0] == "toggle" {
			from := parseCoords(inst[1])
			to := parseCoords(inst[3])
			m.TurnOnGrid(from, to)
			m.TurnOnGrid(from, to)
		}
	}
	fmt.Println(m.TotalBrightness())
}

func main() {
	part1()
	part2()
}
