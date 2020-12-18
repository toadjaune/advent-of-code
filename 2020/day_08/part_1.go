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

func run_for_a_loop(program []ProgramLine) int {
	current_index := 0
	accumulator := 0
	for {
		if program[current_index].executed {
			// We've already executed this specific instruction, let's get out of here
			return accumulator
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

func main() {

	// lines := read_input("input_short")
	lines := read_input("input")

	fmt.Println(run_for_a_loop(parse_program(lines)))

}
