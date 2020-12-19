package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func count_steps(array []int) (int, int) {
	steps_1 := 0
	steps_3 := 0

	for i := 0; i < len(array)-1; i++ {
		if array[i+1]-array[i] == 1 {
			steps_1++
		}
		if array[i+1]-array[i] == 3 {
			steps_3++
		}
	}
	return steps_1, steps_3
}

func main() {

	// lines := read_input("input_short")
	lines := read_input("input")

	integers := extract_integers(lines)
	integers = append(integers, 0)
	sort.Ints(integers)
	integers = append(integers, integers[len(integers)-1]+3)
	steps_1, steps_3 := count_steps(integers)

	fmt.Println(steps_1 * steps_3)

}
