package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func walk_slice(values []int) int {

	for index, val1 := range values {
		for _, val2 := range values[index:] {
			if val1+val2 == 2020 {
				return val1 * val2
			}
		}
	}

	// should not happen on a valid input file
	return 0
}

func main() {
	// Open input file
	// file, err := os.Open("short_input.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read it line by line
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Put values in a slice
	var values []int
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		values = append(values, value)
	}

	fmt.Println(walk_slice(values))

}
