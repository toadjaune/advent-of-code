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

// Doesn't work, because it's not just THE first and THE last line, but a few at each extremity
// func generate_possible_seat_ids() map[int]bool {
//   seat_list := make(map[int]bool)
//
//   // The entire time we were just counting in binary, so generating all possible ids (without accounting for non-existing seats)
//   // is just a matter of generating the entire range.
//
//   for i := 0 ; i <  (1 << 10) ; i++ {
//
//     // We won't generate the first and last line
//     row_number := i >> 3
//     if row_number != 0 &&
//       row_number != (1 << 7) -1 {
//
//       seat_list[i] = true
//
//     }
//   }
//   return seat_list
// }

func main() {

	lines := read_input()

	seat_list := make([]bool, 1024)

	for _, line := range lines {
		seat_list[seat_id(parse_seat_position(line))] = true
	}

	for i := 1; i < len(seat_list)-1; i++ {
		if seat_list[i-1] && !seat_list[i] && seat_list[i+1] {
			fmt.Println(i)
		}
	}
}
