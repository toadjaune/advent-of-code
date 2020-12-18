package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read_input() []string {
	// Open input file
	file, err := os.Open("input")
	//file, err := os.Open("input_short")
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

func regroup_lines(input []string) []string {
	output := make([]string, 1)
	for _, line := range input {
		if line == "" {
			// the group is complete, we move over to the next group
			output = append(output, "")
			continue
		}
		output[len(output)-1] += line
	}
	return output
}

func deduplicate_answers(groups []string) (grouped_answers []map[rune]bool) {
	for _, group_string := range groups {
		group_answers := make(map[rune]bool)
		for _, answer := range group_string {
			group_answers[answer] = true
		}
		grouped_answers = append(grouped_answers, group_answers)
	}
	return grouped_answers
}

func count_total_answers(grouped_answers []map[rune]bool) int {
	counter := 0
	for _, group_answer := range grouped_answers {
		counter += len(group_answer)
	}
	return counter
}

func main() {

	lines := read_input()
	lines = regroup_lines(lines)

	answers_by_group := deduplicate_answers(lines)

	fmt.Println(count_total_answers(answers_by_group))

}
