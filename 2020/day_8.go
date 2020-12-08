package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strings"
	"strconv"
)

// line count of the file
//var lineCount = 612

func swap(operations [612][2]string, i int, op string)(instructions [612][2]string) {
	newOperations := operations
	newOp := "jmp"
	if op == "jmp" {
		newOp = "nop"
	}
	newOperations[i][0] = newOp
	return newOperations
}

func partB(operations [612][2]string)(accumulator int) {
	for i := range operations {
		op := operations[i][0]
		if op == "acc" {
			continue
		} else {
			newOperations := swap(operations, i, op)
			count, state := partA(newOperations)
			if state == "done" {
				return count
			}
		}
	}
	return 0
}

func partA(operations [612][2]string)(accumulator int, state string) {
	var opCounter [612]int
	i := 0
	for {
		if i > len(operations)-1 {
			return accumulator, "done"
		}
		opCounter[i] = opCounter[i]+1
		if opCounter[i] == 2 {
			return accumulator, "loops"
		}
		op := operations[i][0]
		arg, _ := strconv.Atoi(operations[i][1])
		if op == "nop" {
			i++
		} else if op == "jmp" {
			i = i+arg
		} else if op == "acc" {
			accumulator = accumulator+arg
			i++
		}
	}
	return 0, "shit broke"
}

func main() {
	file, err := os.Open("day_8_input.txt")
	//file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var instructions [612][2]string
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		operation := strings.Split(line, " ")
		instructions[i][0] = operation[0]
		instructions[i][1] = operation[1]
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	//fmt.Println(instructions)
	countA, stateA := partA(instructions)
	fmt.Println("Part A:", countA, stateA)
	fmt.Println("Part B:", partB(instructions))
}
