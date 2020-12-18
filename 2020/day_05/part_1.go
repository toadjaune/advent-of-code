package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read_input() []string {
	// Open input file
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

func parse_seat_position(seat_code string) (row_id int, column_id int) {

	row_id = 0
	column_id = 0

	// Extract row number
	for i := 0; i < 7; i++ {

		// Shift row_id by one, allowing to insert next byte at the rightmost position
		row_id = row_id << 1
		if string(seat_code[i]) == "B" {
			row_id++
		}
	}

	// Idem for column
	for i := 7; i < 10; i++ {

		// Shift row_id by one, allowing to insert next byte at the rightmost position
		column_id = column_id << 1
		if string(seat_code[i]) == "R" {
			column_id++
		}
	}
	return row_id, column_id

}

func seat_id(row_id int, column_id int) int {
	return 8*row_id + column_id
}

func main() {

	lines := read_input()

	max_id := 0
	for _, line := range lines {
		line_id := seat_id(parse_seat_position(line))
		if max_id < line_id {
			max_id = line_id
		}
	}

	fmt.Println(max_id)

}
