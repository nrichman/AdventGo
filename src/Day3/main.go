package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func main() {
	fmt.Println(problemOne())
	//fmt.Println(problemTwo())
}

func problemOne() int {
	data := utils.ReadFileAsStringSlice("input_1.txt")
	claims := [1000][1000]int{}
	answer := 0

	for _, line := range data {
		line := strings.Split(line, " ")

		startX, _ := strconv.Atoi(strings.Split(line[2], ",")[0])
		startYStr := strings.Split(line[2], ",")[1]
		startY, _ := strconv.Atoi(startYStr[:len(startYStr)-1])

		lengthX, _ := strconv.Atoi(strings.Split(line[3], "x")[0])
		lengthY, _ := strconv.Atoi(strings.Split(line[3], "x")[1])

		fmt.Printf("%d %d\n", startX, lengthX)
		fmt.Printf("%d %d\n", startY, lengthY)

		for i := startX; i < startX+lengthX; i++ {
			for j := startY; j < startY+lengthY; j++ {
				claims[i][j]++
			}
		}
	}

	for i := 0; i < len(claims); i++ {
		for j := 0; j < len(claims[0]); j++ {
			if claims[i][j] >= 2 {
				answer++
			}
		}
	}

	return answer
}
