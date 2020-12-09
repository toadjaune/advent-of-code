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

// So, I'm quite angry at golang
// hacking maps into behaving like sets is all nice and good but once you ask for more advanced
// primitives such as ensemblist operations, it falls completely short.
// Sure, I can reimplement it, but it's not gonna be quite as good as a native implem, plus I can make mistakes.
// Also, generic types would be a nice addition here, to allow an external library to do it efficiently.
// Everything that is possible is quite cumbersome here
// So I'm throwing an implementation together, but it's not gonna be good.
func set_intersection(set1 map[rune]bool, set2 map[rune]bool) map[rune]bool {
	result := make(map[rune]bool)
	for key := range set1 {
		_, key_is_present := set2[key]
		if key_is_present {
			result[key] = true
		}
	}
	return result
}

func parse_one_line(line string) map[rune]bool {
	result := make(map[rune]bool)
	for _, answer := range line {
		result[answer] = true
	}
	return result
}

func parse_all_lines(lines []string) []map[rune]bool {
	result := make([]map[rune]bool, 0)
	for _, line := range lines {
		result = append(result, parse_one_line(line))
	}
	return result
}

func form_groups(answers []map[rune]bool) [][]map[rune]bool {

	groups := make([][]map[rune]bool, 1)

	for _, answer := range answers {
		if len(answer) > 0 {
			// This is an actual line, we want to add it up to current group
			groups[len(groups)-1] = append(groups[len(groups)-1], answer)
		} else {
			// We've finished the group, moving over to the next one
			groups = append(groups, nil)
		}
	}
	return groups
}

func intersect_group_answers(grouped_answers [][]map[rune]bool) []map[rune]bool {
	intersected_group_answers := make([]map[rune]bool, 0)
	for _, group := range grouped_answers {
		intersected_answers, group := group[0], group[1:]
		for _, answer := range group {
			intersected_answers = set_intersection(intersected_answers, answer)
		}
		intersected_group_answers = append(intersected_group_answers, intersected_answers)
	}
	return intersected_group_answers
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
	answers := parse_all_lines(lines)
	grouped_answers := form_groups(answers)
	intersected_group_answers := intersect_group_answers(grouped_answers)
	fmt.Println(count_total_answers(intersected_group_answers))

}
