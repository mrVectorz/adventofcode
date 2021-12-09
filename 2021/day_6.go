package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func fishCycle(fishes []int) []int {
	currentFishes := fishes
	var newFishes int
	for index, fish := range currentFishes {
		if fish == 0 {
			newFishes += 1
			currentFishes[index] = 6
		} else {
			currentFishes[index] = fish - 1
		}
	}
	for newFishes > 0 {
		currentFishes = append(currentFishes, 8)
		newFishes = newFishes - 1
	}
	return currentFishes
}

func smartIshFishCycle(school [9]int) [9]int {
	newFishes := school[0]
	for i, d := range school[1:] {
		school[i] = d
	}
	school[8] = newFishes
	school[6] = school[6] + newFishes
	return school
}

func main() {
	file, err := os.Open("day_6_input.txt")
	//file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fishes []int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, fish := range line {
			f, _ := strconv.Atoi(fish)
			fishes = append(fishes, f)
		}
	}

	var school [9]int
	for _, fish := range fishes {
		school[fish] += 1
	}

	for day := 1; day <= 80; day++ {
		fishes = fishCycle(fishes)
		//fmt.Println("After", day, "days:", len(fishes))
	}
	fmt.Println(len(fishes))

	//fmt.Println(school)
	partB := 0
	for day := 1; day <= 256; day++ {
		total := 0
		school = smartIshFishCycle(school)
		for _, f := range school {
			total = total + f
		}
		partB = total
		//fmt.Println("After", day, "days:", total)
	}
	fmt.Println("Part B:", partB)
}
