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

// type NumberPaths struct {
//   jolt int
//   paths int
// }
//
// func transform_to_paths(integers []int) (paths []NumberPaths) {
//
//   for _, i := range integers {
//     paths = append(paths, NumberPaths{jolt: i})
//   }
//   return paths
//
// }
//
// func count_possible_paths(paths *[]NumberPaths) {
//   paths
// }

func jolt_to_index(input []int) []int {
	output := make([]int, input[len(input)-1]+1)

	for _, jolt := range input {
		output[jolt] = -1
	}

	return output
}

func count_possible_paths(paths_array []int) []int {

	paths_array[len(paths_array)-1] = 1

	for i := len(paths_array) - 4; 0 <= i; i-- {
		if paths_array[i] != -1 {
			// We don't have such a charger
			continue
		}
		paths_array[i] = paths_array[i+1] + paths_array[i+2] + paths_array[i+3]
	}
	return paths_array
}

func main() {

	// lines := read_input("input_short")
	lines := read_input("input")

	integers := extract_integers(lines)
	integers = append(integers, 0)
	sort.Ints(integers)
	integers = append(integers, integers[len(integers)-1]+3)

	paths := jolt_to_index(integers)

	counted_paths := count_possible_paths(paths)
	fmt.Println(counted_paths[0])

}
