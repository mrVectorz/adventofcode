package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
	"sort"
)

func partA(numbers []int)(a int){
	one:=0
	three:=1
	old:=0
	for _, i := range numbers{
		if i-old == 1 {
			one++
		} else {
			three++
		}
		old=i
	}
	return one*three
}

func slide(w [3]int, n int)(x [3]int){
	var z = [...]int{w[1], w[2], n}
	return z
}

func partB(numbers []int)(b int){
	nums := []int{0,0,0}
	nums = append(nums, numbers...)
	var old = [...]int{0,0,1}
	for i := range nums {
		possibilities := 0
		if i < 3 { continue }
		v := nums[i]-3
		if nums[i] == nums[i-1] {
			continue
		} else {
		if nums[i-3] >= v {
			possibilities = possibilities+old[0]
		}
		if nums[i-2] >= v {
			possibilities = possibilities+old[1]
		}
		if nums[i-1] >= v {
			possibilities = possibilities+old[2]
		}
		//fmt.Println(i, "old", old, "nums:", nums[i-3:i], "nums[i]", nums[i],"v:", v, "possi:", possibilities)
		old = slide(old, possibilities)
		}
	}
	return old[2]
}

func main() {
	file, err := os.Open("day_10_input.txt")
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
	sort.Ints(numbers)
	fmt.Println("Part A:", partA(numbers))
	fmt.Println("Part B:", partB(numbers))
}
