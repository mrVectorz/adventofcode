package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"log"
	"strconv"
)

func main() {
	/*
	test := "my string yes"
	fmt.Println(test)
	s_test := strings.Split(test, " ")
	fmt.Println(s_test[1])
	*/
	file, err := os.Open("day_2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	goodPasswords := 0
	partTwo := 0

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		line := scanner.Text()
		s_line := strings.Split(line, " ")

		count := strings.Split(s_line[0], "-")
		minCount, err := strconv.Atoi(count[0])
		maxCount, err := strconv.Atoi(count[1])
		if err != nil {
			log.Fatal(err)
		}

		letter := s_line[1]
		password := s_line[2]

		total := 0
		for _, value := range password {
			if string(value) == letter[0:1] {
				total++
			}
		}

		if minCount <= total && total <= maxCount {
			goodPasswords++
			//fmt.Printf("%d %d\t%s\t%s\t count:%d\n", minCount, maxCount, letter[0:1], password, total)
		}

		// Part 2
		//fmt.Printf("%d\t%s\t%s\n", minCount, password[minCount-1:minCount], password)
		if password[minCount-1:minCount] == letter[0:1] && password[maxCount-1:maxCount] != letter[0:1] {
			partTwo++
		} else if password[minCount-1:minCount] != letter[0:1] && password[maxCount-1:maxCount] == letter[0:1] {
			partTwo++
		}
	}

	fmt.Println("\nPart 1 - Old good password count:", goodPasswords)
	fmt.Println("Part 2 - Good password count:", partTwo)
}
