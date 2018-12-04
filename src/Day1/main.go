package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func main() {
	ProblemOne()
	ProblemTwo()
}

func ProblemOne() {
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
	fmt.Println(total)
}

func ProblemTwo() {
}
