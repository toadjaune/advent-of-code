package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func read_input() []string {
	// Open input file
	file, err := os.Open("input_short_correct")
	//file, err := os.Open("input")
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
		byr, has_byr := passport["byr"]
		iyr, has_iyr := passport["iyr"]
		eyr, has_eyr := passport["eyr"]
		hgt, has_hgt := passport["hgt"]
		hcl, has_hcl := passport["hcl"]
		ecl, has_ecl := passport["ecl"]
		pid, has_pid := passport["pid"]

		if has_byr && has_iyr && has_eyr && has_hgt && has_hcl && has_ecl && has_pid {
			// All required fields are present, we move on to field structure validation

			byr_struct_ok, _ := regexp.MatchString(`\d{4}`, byr)
			iyr_struct_ok, _ := regexp.MatchString(`\d{4}`, iyr)
			eyr_struct_ok, _ := regexp.MatchString(`\d{4}`, eyr)
			hgt_struct_ok, _ := regexp.MatchString(`(\d{3}cm)|(\d{2}in)`, hgt)
			hcl_struct_ok, _ := regexp.MatchString(`#[0-9a-f]{6}`, hcl)
			ecl_struct_ok, _ := regexp.MatchString(`[a-z]{3}`, ecl)
			pid_struct_ok, _ := regexp.MatchString(`\d{9}`, pid)

			if byr_struct_ok && iyr_struct_ok && eyr_struct_ok && hgt_struct_ok && hcl_struct_ok && ecl_struct_ok && pid_struct_ok {
				// All fields have correct structure, we move on to value validation

				byr_int, _ := strconv.Atoi(byr)
				iyr_int, _ := strconv.Atoi(iyr)
				eyr_int, _ := strconv.Atoi(eyr)
				hgt_ok := false
				if len(hgt) == 5 { // cm
					hgt_cm, _ := strconv.Atoi(hgt[0:3])
					hgt_ok = 150 <= hgt_cm && hgt_cm <= 193
				} else { // in
					hgt_in, _ := strconv.Atoi(hgt[0:2])
					hgt_ok = 59 <= hgt_in && hgt_in <= 76
				}
				// We'd like a set, a map with a dummy value field is the closest we have
				ecl_legal_values := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
				_, ecl_ok := ecl_legal_values[ecl]

				if 1920 <= byr_int && byr_int <= 2002 &&
					2010 <= iyr_int && iyr_int <= 2020 &&
					2020 <= eyr_int && eyr_int <= 2030 &&
					hgt_ok &&
					ecl_ok {

					count++
				}
			}
		}
	}
	return count
}

func main() {

	lines := read_input()

	passports := parse_passports(lines)

	fmt.Println(count_valid_passports(passports))

}
