package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sumTwoCheck(c int, items []int)(int) {
	for _, s := range items {
		if s+c == 2020 {
			fmt.Println(s, " ", c)
			return s
		}
	}
	return 0
}

func main() {
	//file, err := os.Open("./file.txt")
	file, err := os.Open("day_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var items []int

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		item, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}

	printSlice(items)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i:=0; i < len(items); i++ {
		//fmt.Println(i, "\t", items[i])
		b := sumTwoCheck(items[i], items[i:len(items)])
		if b > 0 {
			fmt.Println("part 1: ", items[i]*b)
		}
		for _, s := range items[i:len(items)] {
			c := sumTwoCheck(items[i]+s, items[i:len(items)])
			if c > 0 {
				fmt.Println("part 2: ", items[i]*s*c)
			}
		}
	}
}
