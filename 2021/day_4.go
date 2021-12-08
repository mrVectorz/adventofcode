package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type winner struct {
	card       int
	count      int
	lastNumber int
	line       []int
}

func Split(r rune) bool {
	return r == ',' || r == ' '
}

func isInSlice(num int, intSlice []int) int {
	//return where in the slice
	var found int
	for i, n := range intSlice {
		if num == n {
			found = i
		}
	}
	return found
}

func cardAdder(card [][]int, calledNums []int) int {
	sum := 0
	for _, row := range card {
		for _, i := range row {
			n := isInSlice(i, calledNums)
			if i != calledNums[n] {
				sum = sum + i
			}
		}
	}
	return sum
}

func checkVertWin(numbers []int, cards [][][]int) winner {
	//fastestWinner := winner{0, 99, 0, []int{0, 0, 0, 0, 0}}
	fastestWinner := winner{0, 0, 0, []int{0, 0, 0, 0, 0}}
	//lastWinner := winner{0, 99, 0, []int{0, 0, 0, 0, 0}}
	lastWinner := winner{0, 0, 0, []int{0, 0, 0, 0, 0}}

	for cardIndex, card := range cards {
		rowIndex := 0
		for 5 > rowIndex {
			counter := 0
			var virticalLine []int
			for numberCount, num := range numbers {
				for _, row := range card {
					if row[rowIndex] == num {
						counter += 1
						virticalLine = append(virticalLine, row[rowIndex])
					}
					if counter == 5 {
						fmt.Printf("Bingo! Card %d Column %d After %d number draws\trow %v\n", cardIndex, rowIndex, numberCount+1, row)
						fmt.Println("virticalLine", virticalLine)
						lastWinner = winner{cardIndex, numberCount + 1, num, virticalLine}
						break
					}
				}
				//flip sign tp find fastest winner
				if lastWinner.count > fastestWinner.count {
					fastestWinner = lastWinner
				}

			}
			virticalLine = nil
			rowIndex += 1
		}
	}
	return fastestWinner
}

func checkHorzWin(numbers []int, cards [][][]int) []int {
	// card, call count, row, last number called, highlightSum
	//fastestWinner := []int{0, 99, 0, 0, 0}
	fastestWinner := []int{0, 0, 0, 0, 0}
	//lastWinner := []int{0, 99, 0, 0, 0}
	lastWinner := []int{0, 0, 0, 0, 0}

	for cardIndex, card := range cards {
		highlightSum := 0
		var cardWins []int
		for rowIndex, row := range card {
			//fmt.Println("\nValidating card's", cardIndex, "row", rowIndex, row)
			counter := 0
			//otherCounter := 0
			for numberCount, num := range numbers {
				//fmt.Println("Searching for", num, "in row", row, sort.SearchInts(row, num))
				n := isInSlice(num, row)
				//n := sort.SearchInts(row, num)
				if n < len(row) && row[n] == num {
					fmt.Println("Card:", cardIndex, "found", num, "in row", row)
					highlightSum = highlightSum + num
					counter += 1
				}
				/*
					if counter < 5 {
						otherCounter += 1
					}
				*/
				if counter == 5 {
					fmt.Printf("Bingo! Card %d Row %d After %d number draws\trow %v\n", cardIndex, rowIndex, numberCount+1, row)
					lastWinner = []int{cardIndex, numberCount + 1, rowIndex, num, highlightSum}
					cardWins = append(cardWins, numberCount+1)
					//lastWinner = []int{cardIndex, otherCounter + 1, rowIndex, num, highlightSum}
					break
				}
			}
			//flip sign tp find fastest winner
			if lastWinner[1] > fastestWinner[1] {
				fastestWinner = lastWinner
			}
		}
		fmt.Println("\n\n", cardIndex, cardWins, "\n\n")
		highlightSum = 0
	}
	return fastestWinner
}

func lineToNumbers(characters []string) []int {
	var intSlice []int
	for _, number := range characters {
		number, _ := strconv.Atoi(number)
		intSlice = append(intSlice, number)
	}
	return intSlice
}

func main() {
	//file, err := os.Open("day_4_input.txt")
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var card [][]int
	var cards [][][]int
	var randomNumbers []int
	cardCount := 0

	for scanner.Scan() {
		line := strings.FieldsFunc(scanner.Text(), Split)
		//fmt.Println(line)
		if len(line) > 5 {
			randomNumbers = lineToNumbers(line)
			//fmt.Println("First line", randomNumbers)
		} else if len(line) > 1 {
			card = append(card, lineToNumbers(line))
		} else if len(line) == 0 {
			if len(card) == 0 {
				continue
			}
			//fmt.Println("Empty line", line)
			//fmt.Println("Empty line", card)
			cards = append(cards, card)
			card = nil
			cardCount += 1
			//fmt.Println("Card count:", cardCount)
		}
	}
	// There is no empty line after the last card
	cards = append(cards, card)
	card = nil
	cardCount += 1

	horizBingo := checkHorzWin(randomNumbers, cards)
	virtBingo := checkVertWin(randomNumbers, cards)
	fmt.Println("\nPicks", randomNumbers)
	//fmt.Println("Vertical:", virtBingo, cards[virtBingo[0]])
	fmt.Println("Vertical:", virtBingo, virtBingo.line)
	fmt.Println("Horizontal:", horizBingo, cards[horizBingo[0]][horizBingo[2]])
	//flip sign to win
	if horizBingo[1] > virtBingo.count {
		fmt.Println("H")
		n := isInSlice(horizBingo[3], randomNumbers)
		//fmt.Println(randomNumbers[:n+1])
		cardSum := cardAdder(cards[horizBingo[0]], randomNumbers[:n+1])
		fmt.Println(cardSum * horizBingo[3])
	} else if virtBingo.count > horizBingo[1] {
		fmt.Println("V")
	}
}
