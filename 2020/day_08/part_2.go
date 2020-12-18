package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

type ProgramLine struct {
	executed    bool
	instruction string
	argument    int
}

func parse_program_line(line string) (program_line ProgramLine) {
	fields := strings.Fields(line)
	program_line.instruction = fields[0]
	program_line.argument, _ = strconv.Atoi(fields[1])
	return program_line
}

func parse_program(lines []string) []ProgramLine {
	program := []ProgramLine{}
	for _, line := range lines {
		program = append(program, parse_program_line(line))
	}
	return program
}

func run_program(program []ProgramLine) (bool, int) {
	current_index := 0
	accumulator := 0
	for {
		if current_index == len(program) {
			// We're about to execute the line just after the end of the program, this program doesn't loop
			return true, accumulator
		}
		if program[current_index].executed {
			// We've already executed this specific instruction, this is an infinite loop
			return false, 0
		}
		program[current_index].executed = true

		switch program[current_index].instruction {

		case "nop":
			current_index++
		case "acc":
			accumulator += program[current_index].argument
			current_index++
		case "jmp":
			current_index += program[current_index].argument

		}

	}

}

func find_corrupted_line(program []ProgramLine) int {

	program_copy := make([]ProgramLine, len(program))

	for line_index := 0; line_index < len(program); line_index++ {

		// Refresh the program copy to a clean state
		copy(program_copy, program)

		switch program_copy[line_index].instruction {

		case "nop":
			program_copy[line_index].instruction = "jmp"
		case "jmp":
			program_copy[line_index].instruction = "nop"
		default:
			continue
		}

		valid, accumulator := run_program(program_copy)
		if valid {
			return accumulator
		}
	}

	// Should not be possible
	return -1

}

func main() {

	// lines := read_input("input_short")
	lines := read_input("input")

	fmt.Println(find_corrupted_line(parse_program(lines)))

}
