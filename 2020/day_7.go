package main
// clear plum bags contain 3 dark red bags, 3 dim gold bags, 2 posh black bags, 5 plaid beige bags.

import (
  "os"
  "fmt"
  "log"
  "bufio"
	"strings"
	"regexp"
	"strconv"
)

var partA []string
var rules = make(map[string]string)

func containsString(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func containsColor(color string){
	for i, v := range rules {
		matched, _ := regexp.MatchString(color, v)
		if matched && ! containsString(partA, i) {
			//fmt.Println(i, v)
			partA = append(partA, i)
			containsColor(i)
		}
	}
}

func colorCost(color string)(totalCost int) {
	if color == "no other" {
		return 0
	}
	totalCost = 0
	multiplier := 1
	//fmt.Println(rules[color])
	bags := strings.Split(rules[color], ", ")
	for i := range bags {
		//fmt.Println(bags[i])
		multiplier, _ = strconv.Atoi(bags[i][0:1])
		re := regexp.MustCompile(`[a-z]+ [a-z]+`)
		//fmt.Println(re.FindString(bags[i]))
		//fmt.Println(multiplier)
		cost := colorCost(re.FindString(bags[i]))
		//fmt.Println(totalCost, multiplier, cost)
		totalCost = totalCost+(multiplier*cost)
	}
	return totalCost+1
}

func main() {
  file, err := os.Open("day_7_input.txt")
  //file, err := os.Open("test.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)


  for scanner.Scan() {
    line := scanner.Text()
		//fmt.Printf("%q\n", strings.Split(line, " bags contain "))
		rule := strings.Split(line, " bags contain ")
		rules[rule[0]] = rule[1]
	}
	//Part A
	containsColor("shiny gold")
	//fmt.Println(partA)
	fmt.Println("Part A:", len(partA))

	//Part B
	fmt.Println("Part B:", colorCost("shiny gold")-1)
}
