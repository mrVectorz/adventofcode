package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"log"
	"strconv"
	"regexp"
)

func checkPartA(id []string) (isGood bool){
	//fmt.Println(id)
	for i := range id[0:len(id)-1] {
		if id[i] == "" {
			return false
		}
	}
	return true
}

func checkPartB(id []string) (isGood bool){
	if ! checkPartA(id) {
		return false
	}
	// validate bd
	byr, _ := strconv.Atoi(id[0])
	if byr < 1920 || byr > 2002 {
		return false
	}

	// issue year
	iyr, _ := strconv.Atoi(id[1])
	if iyr < 2010 || iyr > 2020 {
    return false
  }

	// Exp year
	eyr, _ := strconv.Atoi(id[2])
  if eyr < 2020 || eyr > 2030 {
    return false
  }

	// height
	hgtUnit := id[3][len(id[3])-2:len(id[3])]
	hgt, _ := strconv.Atoi(id[3][0:len(id[3])-2])
	if hgtUnit == "cm" && (hgt < 150 || hgt > 193) {
		return false
	} else if hgtUnit == "in" && (hgt < 59 || hgt > 76) {
		return false
	} else if hgtUnit != "in" && hgtUnit != "cm" {
		return false
	}

	// hair colour
	hcl := id[4][1:len(id[4])]
	hclFilter, _ := regexp.MatchString("^[a-z0-9]+$", hcl)
	if id[4][0:1] != "#" {
		return false
	} else if len(hcl) != 6 {
		return false
	} else if ! hclFilter {
		return false
	}

	// eye colour
	eclFilter, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", id[5])
	if ! eclFilter {
		return false
	}

	// passport id
	pidFilter, _ := regexp.MatchString("^[0-9]+$", id[6])
	if len(id[6]) != 9 {
		return false
	} else if ! pidFilter {
		return false
	}
	return true
}

func main() {
	file, err := os.Open("day_4_input.txt")
	//file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Birth Year, Issue Year, Expiration Year, Height, Hair Color, Eye Color, Passport ID, Country ID
	//var byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
	var (
		byr = ""
		iyr = ""
		eyr = ""
		hgt = ""
		hcl = ""
		ecl = ""
		pid = ""
		cid = ""
	)
	var id = [...]string{byr, iyr, eyr, hgt, hcl, ecl, pid, cid}
	var records [298][len(id)]string

	spaceCount := 0

	// build array of all IDs
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		line := scanner.Text()

		// when there's an empty line reset id
		if line == "\n" || line == "" {
			//records = append(records, id)
			records[spaceCount] = id
			//fmt.Printf("%d ID: %s\n", spaceCount, id)
			for i := range id {
				id[i] = ""
			}
			spaceCount++
			continue
		}

		//fmt.Println("DEBUG0:", "spaceCount", spaceCount, line)
		s_line := strings.Split(line, " ")
		for _, value := range s_line {
			//fmt.Println("DEBUG1:", value)
			switch {
			case "byr" == value[0:3]:
				byr = value[4:len(value)]
				id[0] = byr
			case "iyr" == value[0:3]:
				iyr = value[4:len(value)]
				id[1] = iyr
      case "eyr" == value[0:3]:
        eyr = value[4:len(value)]
				id[2] = eyr
      case "hgt" == value[0:3]:
        hgt = value[4:len(value)]
				id[3] = hgt
      case "hcl" == value[0:3]:
        hcl = value[4:len(value)]
				id[4] = hcl
      case "ecl" == value[0:3]:
        ecl = value[4:len(value)]
				id[5] = ecl
      case "pid" == value[0:3]:
        pid = value[4:len(value)]
				id[6] = pid
      case "cid" == value[0:3]:
        cid = value[4:len(value)]
				id[7] = cid
			default:
				fmt.Printf("ERROR VALUE NOT MATCHED -> %s %s\n", value, value[0:3])
			}
		}
	}
	partA := 1
	for i := range records {
		if checkPartA(records[i][:]) {
			partA++
		}
	}
	fmt.Println("Part a:", partA)

	partB := 1
  for i := range records {
    if checkPartB(records[i][:]) {
      partB++
    }
  }
  fmt.Println("Part b:", partB)
}
