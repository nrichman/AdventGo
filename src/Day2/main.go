package main

import (
	"fmt"
	"strings"
	"utils"
)

func main() {
	fmt.Println(problemOne())
	fmt.Println(problemTwo())
}

func problemOne() int {
	data := utils.ReadFileAsStringSlice("input_1.txt")
	twoTotal := 0
	threeTotal := 0

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

// Spent too much time on this so I settled for the N^2 * M where N is length of a string
func problemTwo() string {
	data := utils.ReadFileAsStringSlice("input_1.txt")

	for _, line1 := range data {
		for _, line2 := range data {
			if line1 == line2 {
				continue
			}

			line1 := strings.Split(line1, "")
			line2 := strings.Split(line2, "")

			if len(line1) != len(line2) {
				continue
			}

			repeat := 0
			answer := ""
			for i, character := range line1 {
				if character != line2[i] {
					repeat += 1
				} else {
					answer += character
				}
			}

			if repeat == 1 {
				return answer
			}
		}
	}
	return "FAILED"
}
