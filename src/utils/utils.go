package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFileAsStringSlice(filename string) []string {
	var output []string
	data, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}
