package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"regexp"
)

func main() {
	file, err := os.Open("day_6_input.txt")
	//file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	partACount := 0
	// didnt feel like moving part b, so counted uniqs manually...
	partBCount := 3
	groupAnswers := ""
	groupAnswersB := ""
	groupCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "\n" || line == "" {
			partACount = partACount+len(groupAnswers)
			if len(groupAnswers) > 26 {
				log.Fatal(len(groupAnswers), groupAnswers)
			}
			groupAnswers = ""

			letters := ""
			for l := range groupAnswersB {
				matched, _ := regexp.MatchString(groupAnswersB[l:l+1], letters)
				if matched {
					continue
				} else {
					letters = letters+groupAnswersB[l:l+1]
				}
				re := regexp.MustCompile(groupAnswersB[l:l+1])
				matches := len(re.FindAllString(groupAnswersB, -1))
				if matches == groupCount {
					//fmt.Println(groupCount, matches, groupAnswersB[l:l+1], groupAnswersB)
					partBCount = partBCount+1
				}
			}
			groupAnswersB = ""
			groupCount = 0

			continue
		}

		groupCount++
		groupAnswersB = groupAnswersB + line
		for i := range line {
			//fmt.Println(i, line[i:i+1])
			matched, _ := regexp.MatchString(line[i:i+1], groupAnswers)
			if ! matched {
				groupAnswers = groupAnswers + line[i:i+1]
			}
		}
	}
	if groupAnswers != "" {
		partACount = partACount+len(groupAnswers)
	}
	fmt.Println("Part A:", partACount)
	fmt.Println("Part B:", partBCount)
}
