package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func mean(numbers []int) float64 {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return math.Round(float64(total) / float64(len(numbers)))
}

func fuelCost(positions []int, destination int) int {
	cost := 0
	for _, position := range positions {
		travel := 0
		if position > destination {
			travel = position - destination
		} else if position < destination {
			travel = destination - position
		}
		cost = cost + travel
		//fmt.Println("Moce from", position, "to", destination, ":", travel, "fuel")
	}
	return cost
}

func fancyFuelCost(positions []int, destination int) int {
	cost := 0
	for _, position := range positions {
		travel := 0
		if position > destination {
			travel = position - destination
		} else if position < destination {
			travel = destination - position
		}
		counter := 0
		for i := 1; i <= travel; i++ {
			counter = counter + i
		}
		cost = cost + counter
		//fmt.Println("Moce from", position, "to", destination, ":", travel, "fuel")
	}
	return cost
}

func main() {
	file, err := os.Open("day_7_input.txt")
	//file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var crabs []int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, crab := range line {
			f, _ := strconv.Atoi(crab)
			crabs = append(crabs, f)
		}
	}
	//fmt.Println(crabs)
	//fmt.Println(mean(crabs))
	//fmt.Println(fuelCost(crabs, 2))
	var costs [][]int
	for i, _ := range crabs {
		costs = append(costs, []int{i, fuelCost(crabs, i)})
	}
	//fmt.Println(costs)
	lowestCost := []int{99999999, 99999999}
	for _, cost := range costs {
		if cost[1] < lowestCost[1] {
			lowestCost = cost
		}
	}
	fmt.Println("Part A - position:", lowestCost[0], "fuel cost:", lowestCost[1])

	//fmt.Println(fancyFuelCost(crabs, 5))
	var newCosts [][]int
	for x, _ := range crabs {
		newCosts = append(newCosts, []int{x, fancyFuelCost(crabs, x)})
	}
	lowestCostB := newCosts[0]
	for _, newCost := range newCosts {
		if newCost[1] < lowestCostB[1] {
			lowestCostB = newCost
		}
	}
	//fmt.Println(newCosts)
	fmt.Println("Part B - position:", lowestCostB[0], "fuel cost:", lowestCostB[1])
}
