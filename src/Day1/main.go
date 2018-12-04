package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func main() {
	fmt.Println(ProblemOne())
	fmt.Println(ProblemTwo())
}

func ProblemOne() int {
	total := 0
	data := utils.ReadFileAsStringSlice("input_1.txt")

	for _, line := range data {
		line := strings.Split(line, "")
		value, err := strconv.Atoi(strings.Join(line[1:], ""))
		if err != nil {
			continue
		}

		if line[0] == "+" {
			total += value
		}

		if line[0] == "-" {
			total -= value
		}
	}
	return total
}

func ProblemTwo() int {
	total := 0
	data := utils.ReadFileAsStringSlice("input_1.txt")
	dict := map[int]int{}

	loop := true
	for loop {
		for _, line := range data {
			line := strings.Split(line, "")
			value, err := strconv.Atoi(strings.Join(line[1:], ""))
			if err != nil {
				continue
			}

			if line[0] == "+" {
				total += value
			}

			if line[0] == "-" {
				total -= value
			}

			if _, ok := dict[total]; ok {
				loop = false
				break
			} else {
				dict[total] = 0
			}
		}
	}
	return total
}
