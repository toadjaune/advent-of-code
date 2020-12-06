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

		pos_1, _ := strconv.Atoi(strings.Split(fields[0], "-")[0])
		pos_2, _ := strconv.Atoi(strings.Split(fields[0], "-")[1])
		character := fields[1][0]
		password := fields[2]

    pos_1_correct_char := password[pos_1 - 1] == character
    pos_2_correct_char := password[pos_2 - 1] == character

    // Go has no native XOR operator
		if pos_1_correct_char != pos_2_correct_char {
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
