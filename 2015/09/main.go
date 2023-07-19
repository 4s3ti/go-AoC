package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type distanceMap map[string]map[string]int


func instructions() distanceMap {
	routes := make(distanceMap)

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		instruction := scanner.Text()
		from := strings.Split(instruction, " ")[0]
		to := strings.Split(instruction, " ")[2]
		distance, _ := strconv.Atoi(strings.Split(instruction, " ")[4])

		if routes[from] == nil {
			routes[from] = make(map[string]int)
		}

		if routes[to] == nil {
			routes[to] = make(map[string]int)
		}

		routes[from][to] = distance
		routes[to][from] = distance
	}



	return routes
}

func isInSlice(slice []string, item string) bool {
	for _, sliceItem := range slice {
		if sliceItem == item {
			return true
		}
	}
	return false
}

func cities(distanceMap distanceMap) []string {
	cities := []string{}
	for city := range distanceMap {
		cities = append(cities, city)
	}

	return cities
}

func makeRoutes(cities []string) [][]string {
	if len(cities) == 0 {
		return [][]string{}
	} else if len(cities) == 1 {
		return [][]string{cities}
	}

	routes := [][]string{}


	for i, city := range cities {
		remainingCities := make([]string, len(cities)-1)
		copy(remainingCities[:i], cities[:i])
		copy(remainingCities[i:], cities[i+1:])

		subRoutes := makeRoutes(remainingCities)


		for _, subRoute := range subRoutes {
			routes = append(routes, append([]string{city}, subRoute...))
			
		}

	}

	return routes
}

func shortestRoute(distanceMap distanceMap) ([]string, int) {
	// routes := validRoutes(distanceMap)
	cities := cities(distanceMap)
	routes := makeRoutes(cities)

	var shortestDistance int
	var shortestRoute []string

	for i, route := range routes {
		distance := 0

		for n := 0; n < len(route)-1; n++ {
			from := route[n]
			to := route[n+1]
			distance += distanceMap[from][to]
		}
		if i == 0 {
			shortestRoute = route
			shortestDistance = distance 
		} else if  distance < shortestDistance {
			shortestDistance = distance
			shortestRoute = route
		}
	}
	return shortestRoute, shortestDistance
}

func longestRoute(distanceMap distanceMap) ([]string, int) {
	// routes := validRoutes(distanceMap)
	cities := cities(distanceMap)
	routes := makeRoutes(cities)

	var longestDistance int
	var longestRoute []string

	for i, route := range routes {
		distance := 0

		for n := 0; n < len(route)-1; n++ {
			from := route[n]
			to := route[n+1]
			distance += distanceMap[from][to]
		}
		if i == 0 {
			longestRoute = route
			longestDistance = distance 
		} else if  distance > longestDistance {
			longestDistance = distance
			longestRoute = route
		}
	}
	return longestRoute, longestDistance
}

func visualizeMap(distanceMap distanceMap) {
	for route, distances := range distanceMap {
		fmt.Printf("%v: ", route)
		for destination, distance := range distances {
			fmt.Printf("%v: %v, ", destination, distance)
		}
		fmt.Println()
	}
}

func main() {
	distanceMap := instructions()
	shortestRoute, shortestDistance := shortestRoute(distanceMap)
	longestRoute, longestDistance := longestRoute(distanceMap)

	fmt.Printf("Part1: The shortest route: %v, distannce: %v\n", shortestRoute, shortestDistance)
	fmt.Printf("Part2: The longest route %v, distance: %v\n", longestRoute, longestDistance)
}
