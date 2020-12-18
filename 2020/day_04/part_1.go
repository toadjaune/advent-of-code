package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	"strings"
)

func read_input() []string {
	// Open input file
	//file, err := os.Open("input_short")
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
	for row_index := 0; row_index < len(rows); row_index = row_index + increment_row {

		row := rows[row_index]
		// 35 is the ASCII code for X
		if row[column%len(row)] == 35 {
			trees++
		}
		column = column + increment_column
	}
	return trees
}

func parse_passports(rows []string) []map[string]string {

	var passports []map[string]string
	current_passport := map[string]string{}

	for _, row := range rows {

		// if the line is empty, the current passport is complete, we flush it and start a new one
		if row == "" {
			passports = append(passports, current_passport)
			current_passport = make(map[string]string)
		} else {
			for _, word := range strings.Fields(row) {
				values := strings.Split(word, ":")
				current_passport[values[0]] = values[1]
			}
		}

	}
	return passports
}

func count_valid_passports(passports []map[string]string) int {
	count := 0
	for _, passport := range passports {
		_, has_byr := passport["byr"]
		_, has_iyr := passport["iyr"]
		_, has_eyr := passport["eyr"]
		_, has_hgt := passport["hgt"]
		_, has_hcl := passport["hcl"]
		_, has_ecl := passport["ecl"]
		_, has_pid := passport["pid"]

		if has_byr && has_iyr && has_eyr && has_hgt && has_hcl && has_ecl && has_pid {
			count++
		}
	}
	return count
}

func main() {

	lines := read_input()

	passports := parse_passports(lines)

	fmt.Println(count_valid_passports(passports))

}
