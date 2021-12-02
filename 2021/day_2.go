package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day_2_input.txt")
	//file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [][]int

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		//fmt.Println(line)
		depth, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println(line, err)
		}
		instruction := 0
		switch line[0] {
		case "forward":
			instruction = 0
		case "up":
			instruction = 1
		case "down":
			instruction = 2
		}
		cmd := []int{instruction, depth}
		lines = append(lines, cmd)
	}

	//fmt.Println(len(lines), lines[0])

	depth := 0
	horizontal := 0
	Depth := 0

	//part 1
	for _, d := range lines {
		//fmt.Println(d[0], d[1])
		if d[0] == 0 {
			horizontal = horizontal + d[1]
			Depth = Depth + (depth * d[1])
		} else if d[0] == 1 {
			depth = depth - d[1]
		} else if d[0] == 2 {
			depth = depth + d[1]
		}
	}
	fmt.Println("Part 1")
	fmt.Println("Depth:", depth, "\nHorizontal:", horizontal, "\nTotal:", horizontal*depth)
	fmt.Println("---\nPart 2")
	fmt.Println("Depth:", Depth*horizontal)
}
