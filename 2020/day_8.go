package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strings"
	"strconv"
)

type instruction struct{
	op string
	arg string
}


func swap(operations []instruction, i int, op string)(instructions []instruction) {
	newOperations := make([]instruction, len(operations)) //when copying the slice needs to already have the right length
	copy(newOperations, operations)
	newOp := "jmp"
	if op == "jmp" {
		newOp = "nop"
	}
	newOperations[i].op = newOp
	return newOperations
}

func partB(operations []instruction)(accumulator int) {
	for i := range operations {
		op := operations[i].op
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

func partA(operations []instruction)(accumulator int, state string) {
	opCounter := make([]int, len(operations))
	// so we can't use `var opCounter [len(operators)]int`
	// and the bellow is slow(ish)
	/*
	var opCounter []int
	//build op array
	for range operations {
		opCounter = append(opCounter, 0)
	}
	*/
	i := 0
	for {
		if i > len(operations)-1 {
			return accumulator, "done"
		}

		opCounter[i] = opCounter[i]+1
		if opCounter[i] == 2 {
			return accumulator, "loops"
		}

		op := operations[i].op
		arg, _ := strconv.Atoi(operations[i].arg)
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

	var instructions []instruction
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		operation := strings.Split(line, " ")
		instructions = append(instructions, instruction{operation[0], operation[1]})
		//instructions[i][0] = operation[0]
		//instructions[i][1] = operation[1]
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
