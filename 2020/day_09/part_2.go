package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func extract_integers(lines []string) (integers []int) {
	for _, line := range lines {
		integer, _ := strconv.Atoi(line)
		integers = append(integers, integer)
	}
	return integers
}

func has_valid_components(dumped_numbers []int, window_length int, index int) bool {
	for i := index - window_length; i < index; i++ {
		for j := i + 1; j < index; j++ {
			if dumped_numbers[i]+dumped_numbers[j] == dumped_numbers[index] {
				return true
			}
		}
	}
	return false
}

func find_without_components(dumped_numbers []int, window_size int) int {

	for i := window_size; i < len(dumped_numbers); i++ {
		if !has_valid_components(dumped_numbers, window_size, i) {
			return dumped_numbers[i]
		}
	}
	panic("yolo")
}

func find_range(dumped_numbers []int, sum int) (int, int) {
	for i := 0; i < len(dumped_numbers); i++ {
		acc := 0
		for j := i; acc < sum; j++ {
			acc += dumped_numbers[j]
			if acc == sum {
				return i, j
			}
		}
	}
	panic("WTF")
}

func lowest_highest(array []int) (int, int) {
	min := array[0]
	max := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
		}
		if max < array[i] {
			max = array[i]
		}
	}
	return min, max
}

func main() {

	// lines := read_input("input_short")
	// window_size := 5
	lines := read_input("input")
	window_size := 25

	integers := extract_integers(lines)

	invalid_number := find_without_components(integers, window_size)

	fmt.Println(invalid_number)

	min_index, max_index := find_range(integers, invalid_number)

	lowest, highest := lowest_highest(integers[min_index : max_index+1])

	fmt.Println(lowest + highest)

}
