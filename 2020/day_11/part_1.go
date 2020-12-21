package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read_input(filename string) []string {
	// Open input file
	file, err := os.Open(filename)
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

// Check if a given seat is occupied
func is_occupied(matrix [][]rune, line int, column int) int {

	if line < 0 || len(matrix)-1 < line || column < 0 || len(matrix[0])-1 < column {
		// Out-of-bounds input coordinates are legal input
		// We just consider there's an empty seat here
		return 0
	}

	if matrix[line][column] == rune("#"[0]) {
		return 1
	}

	return 0

}

func count_neighbors(matrix [][]rune, line int, column int) int {
	return is_occupied(matrix, line-1, column) +
		is_occupied(matrix, line-1, column-1) +
		is_occupied(matrix, line, column-1) +
		is_occupied(matrix, line+1, column-1) +
		is_occupied(matrix, line+1, column) +
		is_occupied(matrix, line+1, column+1) +
		is_occupied(matrix, line, column+1) +
		is_occupied(matrix, line-1, column+1)
}

func compute_next_cell_state(matrix [][]rune, line int, column int) rune {
	switch matrix[line][column] {
	case rune("."[0]):
		// Floor. No possible change.
		return rune("."[0])
	case rune("L"[0]):
		// Empty seat
		if count_neighbors(matrix, line, column) > 0 {
			return rune("L"[0])
		} else {
			return rune("#"[0])
		}
	case rune("#"[0]):
		// Empty seat
		if count_neighbors(matrix, line, column) > 3 {
			return rune("L"[0])
		} else {
			return rune("#"[0])
		}
	}
	panic("WTF")
}

func compute_next_grid_state(input_matrix [][]rune) (output_matrix [][]rune, changed bool) {

	for line_index := range input_matrix {
		output_matrix = append(output_matrix, make([]rune, len(input_matrix[0])))
		for column_index := range input_matrix[0] {

			output_matrix[line_index][column_index] = compute_next_cell_state(input_matrix, line_index, column_index)
			changed = changed || input_matrix[line_index][column_index] != output_matrix[line_index][column_index]

		}

	}

	return output_matrix, changed

}

func convert_to_rune_matrix(lines []string) (matrix [][]rune) {
	for _, line := range lines {
		matrix = append(matrix, []rune(line))
	}
	return matrix
}

func run_until_stable(matrix [][]rune) [][]rune {
	var changed bool
	for {
		matrix, changed = compute_next_grid_state(matrix)
		if !changed {
			return matrix
		}
	}
}

func count_occupied_seats(matrix [][]rune) (count int) {

	for i := range matrix {
		for j := range matrix[0] {
			count += is_occupied(matrix, i, j)
		}
	}
	return count

}

func main() {

	// lines := read_input("input_short")
	lines := read_input("input")

	matrix := convert_to_rune_matrix(lines)
	matrix = run_until_stable(matrix)

	fmt.Println(count_occupied_seats(matrix))

}
