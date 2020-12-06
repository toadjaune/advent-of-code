package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func count_valid(lines []string) int {
	valid_lines_count := 0

	for _, line := range lines {

		fields := strings.Fields(line)

		min_count, _ := strconv.Atoi(strings.Split(fields[0], "-")[0])
		max_count, _ := strconv.Atoi(strings.Split(fields[0], "-")[1])
		character := fields[1][0]
		password := fields[2]

		occurrences_number := strings.Count(password, string(character))

		if min_count <= occurrences_number && occurrences_number <= max_count {
			valid_lines_count++
		}
	}

	return valid_lines_count
}

func main() {
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

	fmt.Println(count_valid(lines))

}
