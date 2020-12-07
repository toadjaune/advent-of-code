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

func count_trees_descent(rows []string, increment_column int, increment_row int) int {

	trees := 0
  column := 0
  for row_index := 0 ; row_index < len(rows) ; row_index = row_index + increment_row {

    row := rows[row_index]
		// 35 is the ASCII code for X
		if row[column % len(row)] == 35 {
			trees++
		}
		column = column + increment_column
	}
	return trees
}

func main() {

	lines := read_input()

	// fmt.Println(count_trees_descent(lines, 1, 1))
	// fmt.Println(count_trees_descent(lines, 3, 1))
	// fmt.Println(count_trees_descent(lines, 5, 1))
	// fmt.Println(count_trees_descent(lines, 7, 1))
	// fmt.Println(count_trees_descent(lines, 1, 2))
	fmt.Println(count_trees_descent(lines, 1, 1) * count_trees_descent(lines, 3, 1) * count_trees_descent(lines, 5, 1) * count_trees_descent(lines, 7, 1) * count_trees_descent(lines, 1, 2) )

}
