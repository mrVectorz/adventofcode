package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
//	"strings"
)

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func treeMapping(vert int, horz int, items []string)(hits int){
	mvtCount := 1
	Treehits := 0
	skip := false
	for index, value := range items {
		if index == 0 && vert > 0 {
			continue
		} else if index == 1 && vert > 1 {
			continue
		}

		if skip == true {
			skip = false
			continue
		} else if skip == false && vert > 1 {
			skip = true
		}
		mvtCount = mvtCount+horz
		if mvtCount > len(value) {
			mvtCount = mvtCount-len(value)
		}
		//fmt.Println(value)
		if value[mvtCount-1:mvtCount] == "#" {
			Treehits++
			//fmt.Printf("%s^ <- Hit!\n", strings.Repeat(" ", mvtCount-1))
		}
		//fmt.Printf("mvtCount: %d\ttreeHit: %d\n", mvtCount, Treehits)
	}
	fmt.Println(Treehits)
	return Treehits
}

func main() {
	//file, err := os.Open("./test.txt")
	file, err := os.Open("day_3_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var items []string

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		item := scanner.Text()
		items = append(items, item)
	}

	//printSlice(items)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//totalHits := 0
	totalHits := treeMapping(1, 1, items)
	fmt.Printf("Total hits: %d\n", totalHits)
	totalHits = totalHits*treeMapping(1, 3, items)
	fmt.Printf("Total hits: %d\n", totalHits)
	totalHits = totalHits*treeMapping(1, 5, items)
	fmt.Printf("Total hits: %d\n", totalHits)
	totalHits = totalHits*treeMapping(1, 7, items)
	fmt.Printf("Total hits: %d\n", totalHits)
	totalHits = totalHits*treeMapping(2, 1, items)
	fmt.Printf("Total hits: %d\n", totalHits)

	/*
	//part a
	mvtCount := 1 //horizontal
	treeHit := 0
	for index, value := range items {
		// we basically skip the first row
		if index == 0 {
			continue
		}
		mvtCount = mvtCount+3
		if mvtCount > len(value) {
			mvtCount = mvtCount-len(value)
		}
		fmt.Println(value)
		if value[mvtCount-1:mvtCount] == "#" {
			treeHit++
			fmt.Printf("%s^ <- Hit!\n", strings.Repeat(" ", mvtCount-1))
		}
		fmt.Printf("mvtCount: %d\ttreeHit: %d\n", mvtCount, treeHit)
	}
	*/
}
