package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	// "strings"
)

func read_input() []string {
	// Open input file
	// file, err := os.Open("input_short")
	file, err := os.Open("input")
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
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func count_trees_descent(rows []string) int {
	increment := 3

	trees := 0
	for row_index, row := range rows {

		// 35 is the ASCII code for X
		if row[(row_index*increment)%len(row)] == 35 {
			trees++
		}
	}
	return trees
}

func main() {

	lines := read_input()

	fmt.Println(count_trees_descent(lines))

}
