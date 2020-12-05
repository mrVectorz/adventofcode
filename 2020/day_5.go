package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"log"
	"strconv"
	"sort"
)


func main() {
	var replacer = strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")

	file, err := os.Open("day_5_input.txt")
	//file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var maxSeatID int64
	var seatIDs []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row, _ := strconv.ParseInt(replacer.Replace(line[0:7]), 2, 0)
		column, _ := strconv.ParseInt(replacer.Replace(line[7:10]), 2, 0)
		seatID := row*8+column
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
		seatIDs = append(seatIDs, int(seatID))
		//fmt.Println(seatID)
		//fmt.Println(line, line[0:6], row, line[7:10], column, seatID)
	}
	// Part A
	fmt.Println("Part A:", maxSeatID)

	// Part B
	sort.Ints(seatIDs)
	temp := seatIDs[0]
	for _, v := range seatIDs {
		if v-temp > 1 {
			fmt.Println(v-1)
			break
		}
		temp = v
	}
	//fmt.Println(seatIDs)
}
