package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type segment struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func sortedInts(a int, b int) [2]int {
	if a < b {
		return [2]int{a, b}
	}
	return [2]int{b, a}
}

func mapGenA(grid [][]int, segments []segment) [][]int {
	for _, segment := range segments {
		if segment.x1 == segment.x2 || segment.y1 == segment.y2 {
			//fmt.Println(segment)
			if segment.x1 != segment.x2 {
				n := sortedInts(segment.x1, segment.x2)
				for i := n[0]; i <= n[1]; i++ {
					grid[segment.y1][i] += 1
				}
			} else if segment.y1 != segment.y2 {
				n := sortedInts(segment.y1, segment.y2)
				for i := n[0]; i <= n[1]; i++ {
					grid[i][segment.x1] += 1
				}
			}
		}
	}
	return grid
}

func isDiag(seg segment) bool {
	nOne := sortedInts(seg.x1, seg.x2)
	nTwo := sortedInts(seg.y1, seg.y2)
	if (nOne[1] - nOne[0]) == (nTwo[1] - nTwo[0]) {
		return true
	}
	return false
}

func mapGenB(grid [][]int, segments []segment) [][]int {
	for _, segment := range segments {
		if segment.x1 == segment.x2 || segment.y1 == segment.y2 {
			//fmt.Println(segment)
			if segment.x1 != segment.x2 {
				n := sortedInts(segment.x1, segment.x2)
				for i := n[0]; i <= n[1]; i++ {
					grid[segment.y1][i] += 1
				}
			} else if segment.y1 != segment.y2 {
				n := sortedInts(segment.y1, segment.y2)
				for i := n[0]; i <= n[1]; i++ {
					grid[i][segment.x1] += 1
				}
			}
			//continue
		} else if isDiag(segment) {
			errorStop := 0
			x := segment.x1
			y := segment.y1
			for x != segment.x2 {
				errorStop += 1
				//fmt.Println("Y", y, "X", x)
				grid[y][x] += 1
				if segment.x1 > segment.x2 {
					x = x - 1
				} else if segment.x1 < segment.x2 {
					x = x + 1
				}
				if segment.y1 > segment.y2 {
					y = y - 1
				} else if segment.y1 < segment.y2 {
					y = y + 1
				}
				if errorStop > 1000 {
					fmt.Println("ERROR", segment, "x")
					break
				}
			}
		} else {
			fmt.Println(segment)
		}
	}
	return grid
}
func main() {
	file, err := os.Open("day_5_input.txt")
	//file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []segment

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		segmentStart := strings.Split(line[0], ",")
		segmentEnd := strings.Split(line[1], ",")
		x1, _ := strconv.Atoi(segmentStart[0])
		y1, _ := strconv.Atoi(segmentStart[1])
		x2, _ := strconv.Atoi(segmentEnd[0])
		y2, _ := strconv.Atoi(segmentEnd[1])
		segment := segment{x1, y1, x2, y2}
		//fmt.Println(line)
		lines = append(lines, segment)
	}

	largestX := 0
	largestY := 0
	for _, segment := range lines {
		switch {
		case segment.x1 > largestX:
			largestX = segment.x1
		case segment.x2 > largestX:
			largestX = segment.x2
		case segment.y1 > largestY:
			largestY = segment.y1
		case segment.y2 > largestY:
			largestY = segment.y2
		}
	}
	//fmt.Println("largestX:", largestX, "largestY:", largestY)

	var grid [][]int
	for i := 0; i < largestY+1; i++ {
		x := make([]int, largestX+1)
		grid = append(grid, x)
	}
	//var grid [largestY + 1][largestX + 1]int

	//fmt.Println("emp:", grid)

	/* PART A
	//fmt.Println("h+v lines", mapGen(grid, lines))
	mapA := mapGenA(grid, lines)
	partA := 0
	for _, y := range mapA {
		for _, x := range y {
			if x >= 2 {
				partA += 1
			}
		}
	}
	fmt.Println("Part A:", partA)
	*/

	//fmt.Println("emp:", grid)
	mapB := mapGenB(grid, lines)
	//fmt.Println("h+v+d lines", mapB)
	partB := 0
	for _, y := range mapB {
		for _, x := range y {
			if x >= 2 {
				partB += 1
			}
		}
	}
	fmt.Println("Part B:", partB)

	//fmt.Println(len(lines), lines[0])
}
