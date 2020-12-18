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

type BagAmount struct {
	amount int
	kind   string
}

func parse_line(line string) (container string, contents []BagAmount) {

	parts := strings.SplitN(line, " bags contain ", 2)
	container = parts[0]

	if parts[1] == "no other bags." {
		// We return an empty contents
		return container, contents
	}

	each_bags := regexp.MustCompile(` bags?(, |.)`).Split(parts[1], -1)

	// There's an empty string at the end of the array
	for _, bag_spec := range each_bags[:len(each_bags)-1] {
		// I was gonna assume that the number is only one digit.
		// Pretty sure it's a trap
		split_bag_spec := strings.SplitN(bag_spec, " ", 2)
		amount, _ := strconv.Atoi(split_bag_spec[0])
		bag_amount := BagAmount{
			amount: amount,
			kind:   split_bag_spec[1],
		}
		contents = append(contents, bag_amount)
	}
	return container, contents
}

// All bag types are in here, eve those whithout valid children
func build_allowed_children(lines []string) map[string][]BagAmount {
	allowed_children := make(map[string][]BagAmount, 0)
	for _, line := range lines {
		container, contents := parse_line(line)
		allowed_children[container] = contents
	}
	return allowed_children
}

// Only bag types with a valid parent are in here
func build_allowed_parents(allowed_children map[string][]BagAmount) map[string][]string {
	allowed_parents := make(map[string][]string)
	for parent, children := range allowed_children {
		for _, child := range children {
			parents, is_present := allowed_parents[child.kind]
			if is_present {
				allowed_parents[child.kind] = append(parents, parent)
			} else {
				allowed_parents[child.kind] = []string{parent}
			}
		}
	}
	return allowed_parents
}

func find_all_recursive_parents(allowed_parents map[string][]string, bag_kind string, accumulator map[string]bool) map[string]bool {

	parents, can_have_parent := allowed_parents[bag_kind]

	if !can_have_parent {
		// The current bag type cannot ever have a parent, that's a structural stop condition
		return accumulator
	}

	for _, parent := range parents {
		_, is_already_known_parent := accumulator[parent]
		if is_already_known_parent {
			// This is our real recursion-breaking stop condition
			continue
		} else {
			// We add this parent to the list of known parents
			accumulator[parent] = true
			// Then we look for its parents
			accumulator = find_all_recursive_parents(allowed_parents, parent, accumulator)
		}
	}

	return accumulator

}

func main() {

	// lines := read_input("input_short")
	lines := read_input("input")

	allowed_children := build_allowed_children(lines)
	allowed_parents := build_allowed_parents(allowed_children)

	my_parents := find_all_recursive_parents(allowed_parents, "shiny gold", map[string]bool{})

	// fmt.Println(allowed_children)
	// fmt.Println(allowed_parents)
	// fmt.Println(my_parents)
	fmt.Println(len(my_parents))

}
