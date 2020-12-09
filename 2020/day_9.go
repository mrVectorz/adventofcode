package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
)

func findPair(window []int, sum int)(noMatch bool) {
	for i, v := range window {
		for _, x := range window[i:26] {
			if v == x {
				continue
			}
			if v+x == sum {
				return true
			}
		}
	}
	return false
}

func partA(numbers []int)(badNum int) {
	for i := range numbers {
		if i < 25 {
			continue
		}
		if findPair(numbers[i-25:i], numbers[i+1]) {
			continue
		} else {
			return numbers[i+1]
		}
	}
	return badNum
}

func sumSlice(ints []int)(int) {
	result := 0
	for _, i := range ints {
		result += i
	}
	return result
}

func minIntSlice(v []int) (m int) {
	for i, e := range v {
		if i==0 || e < m {
			m = e
		}
	}
	return
}
func maxIntSlice(v []int) (m int) {
	for i, e := range v {
		if i==0 || e > m {
			m = e
		}
	}
	return
}

func partB(numbers []int, sum int)(weakness int) {
	for i := 2; i < 200; i++ {
		for n := range numbers {
			value := sumSlice(numbers[n:n+i])
			if value == sum {
				min := minIntSlice(numbers[n:n+i])
				max := maxIntSlice(numbers[n:n+i])
				return min+max
			}
		}
	}
	return 0
}

func main() {
	file, err := os.Open("day_9_input.txt")
	//file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}
	//fmt.Println(numbers)
	invalidNumber := partA(numbers)
	fmt.Println("Part A:", invalidNumber)
	fmt.Println("Part B:", partB(numbers, invalidNumber))
}
