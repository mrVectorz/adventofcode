package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day_1_input.txt")
	//file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []int

	first := true
	prevDepth := 0
	countInc := 0
	countDec := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		depth, _ := strconv.Atoi(line)
		lines = append(lines, depth)
		if first == true {
			first = false
			fmt.Println(line, "First line")
			prevDepth = depth
			continue
		}
		if depth > prevDepth {
			countInc += 1
			//fmt.Println(line, "Increased")
		} else if depth < prevDepth {
			countDec += 1
			//fmt.Println(line, "Decreased")
		}
		prevDepth = depth
	}
	fmt.Println("Decreased:", countDec)
	fmt.Println("Increased:", countInc)

	//fmt.Println(len(lines), lines[0])
	prevSum := 0
	sumInc := 0
	for i, d := range lines {
		if i+2 >= len(lines) {
			break
		}
		sum := d + lines[i+1] + lines[i+2]
		if sum > prevSum {
			sumInc += 1
		}
		prevSum = sum
	}
	// minus one due to the start value
	fmt.Println(sumInc - 1)
}
