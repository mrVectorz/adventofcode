package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func maskGen(flag string, bitList []byte) int {
	zeroCount := bytes.Count(bitList, []byte{'0'})
	oneCount := bytes.Count(bitList, []byte{'1'})
	//fmt.Println("zeroCount", zeroCount, "oneCount", oneCount)
	if zeroCount > oneCount && flag == "o" {
		return 0
	} else if oneCount > zeroCount && flag == "o" {
		return 1
	} else if zeroCount == oneCount && flag == "o" {
		return 1
	} else if zeroCount > oneCount && flag == "c" {
		return 1
	} else if oneCount > zeroCount && flag == "c" {
		return 0
	} else if zeroCount == oneCount && flag == "c" {
		return 0
	}
	fmt.Println("ERROR")
	return 2
}

func partTwo(flag string, bitSlice [][]int) int64 {
	//fmt.Println("bitSlice:", bitSlice)
	filteredSlice := bitSlice
	for len(filteredSlice) > 1 {
		//fmt.Println("filteredSlice:", filteredSlice, "len:", len(filteredSlice))
		for i, _ := range bitSlice[0] {
			if 1 == len(filteredSlice) {
				break
			}
			// get the byte slice for maskGen
			var column []byte
			for _, b := range filteredSlice {
				column = strconv.AppendInt(column, int64(b[i]), 10)
				//column = append(column, b[i])
			}
			mask := maskGen(flag, column)
			//fmt.Println("mask:", mask)
			var goodOnes [][]int
			for _, data := range filteredSlice {
				//fmt.Println("filteredSlice", filteredSlice)
				//fmt.Println("mask:", mask, "mask index:", i, "slice index", index, "data:", data, "compared bit", data[i])
				if data[i] == mask {
					goodOnes = append(goodOnes, data)
				}
				/*
					if data[i] != mask {
						fmt.Println("Removing", data)
						lastIndex := len(filteredSlice) - 1
						//fmt.Println("last index", lastIndex, "len", len(filteredSlice))
						if index >= lastIndex {
							filteredSlice = filteredSlice[:lastIndex]
							continue
						}
						fmt.Println(filteredSlice[:index])
						filteredSlice = append(filteredSlice[:index], filteredSlice[index+1:]...)
					}
				*/
			}
			filteredSlice = goodOnes
		}
	}
	var dec string
	fmt.Println("flag:", flag, "slice", filteredSlice)
	for _, v := range filteredSlice[0] {
		d := strconv.Itoa(v)
		dec = dec + d
	}
	decimal, _ := strconv.ParseInt(dec, 2, 64)
	return decimal
}

func bitSelector(criteria string, bitArray [][]int) int64 {
	// This is a slice that I call an Array to keep confusing me
	filteredArray := bitArray
	Ncriteria := strings.Split(criteria, "")
	//fmt.Println("Array len:", len(filteredArray))
	fmt.Println("Criteria:", criteria)
	for len(filteredArray) > 1 {
		//fmt.Println(len(filteredArray))
		for i, c := range Ncriteria {
			if 1 == len(filteredArray) {
				break
			}
			//fmt.Println(i, c)
			w, _ := strconv.Atoi(c)
			for x, d := range filteredArray {
				//fmt.Println("bit", i, "bitData", c, "slice index", x, "slice data", d, "slice len", len(filteredArray))
				if x >= len(filteredArray) {
					break
				}
				if d[i] != w {
					lastIndex := len(filteredArray) - 1
					if x == lastIndex {
						filteredArray = filteredArray[:lastIndex]
						continue
					}
					filteredArray = append(filteredArray[:x], filteredArray[x+1:]...)
				}
			}
			//fmt.Println(filteredArray)
			//break
		}
	}
	var dec string
	fmt.Println(filteredArray)
	for _, v := range filteredArray[0] {
		d := strconv.Itoa(v)
		dec = dec + d
	}
	decimal, _ := strconv.ParseInt(dec, 2, 64)
	return decimal
}

func main() {
	file, err := os.Open("day_3_input.txt")
	//file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [][]int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		//fmt.Println(line)
		var bitArray []int
		for _, d := range line {
			d, _ := strconv.Atoi(d)
			bitArray = append(bitArray, d)
		}
		lines = append(lines, bitArray)
	}

	var gamma string
	var epsilon string
	for i, _ := range lines[0] {
		//fmt.Println(i)
		//var column []int
		var countZero int
		var countOne int
		for _, v := range lines {
			if v[i] == 0 {
				countZero += 1
			} else if v[i] == 1 {
				countOne += 1
			}
			//column = append(column, v)
		}
		if countZero > countOne {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		}
	}
	decimalEpsilon, _ := strconv.ParseInt(epsilon, 2, 64)
	decimalGamma, _ := strconv.ParseInt(gamma, 2, 64)
	fmt.Println("epsilon:", decimalEpsilon)
	fmt.Println("gamma:", decimalGamma)
	fmt.Println("total:", decimalEpsilon*decimalGamma)

	//oxyRating := bitSelector(gamma, lines)
	//co2Rating := bitSelector(epsilon, lines)
	//fmt.Println(oxyRating)
	//fmt.Println(co2Rating)

	//fmt.Println("test1:", maskGen("o", []byte{'1', '0'}))
	//fmt.Println("lines ox", lines)
	oxyRatingTwo := partTwo("o", lines)
	//fmt.Println("lines co2", lines)
	co2RatingTwo := partTwo("c", lines)
	fmt.Println(oxyRatingTwo, "*", co2RatingTwo, "=", oxyRatingTwo*co2RatingTwo)
}
