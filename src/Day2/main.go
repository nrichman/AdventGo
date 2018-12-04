package main

import (
	"fmt"
	"strings"
	"utils"
)

func main() {
	fmt.Println(ProblemOne())
}

func ProblemOne() int {
	twoTotal := 0
	threeTotal := 0
	data := utils.ReadFileAsStringSlice("input_1.txt")

	for _, line := range data {
		wordmap := map[string]int{}

		line := strings.Split(line, "")
		for _, character := range line {
			wordmap[character] += 1
		}

		two := false
		three := false

		for _, v := range wordmap {
			if v == 2 && !two {
				twoTotal += 1
				two = true
			}

			if v == 3 && !three {
				threeTotal += 1
				three = true
			}
		}
	}
	return twoTotal * threeTotal
}
