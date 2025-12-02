package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ValueRange struct {
	lower int
	upper int
}

func read_and_parse_input() []ValueRange {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	ranges := make([]ValueRange, 0)
	for _, s := range strings.Split(string(input), ",") {
		splt := strings.Split(s, "-")

		lower, _ := strconv.Atoi(strings.TrimSpace(splt[0]))
		upper, _ := strconv.Atoi(strings.TrimSpace(splt[1]))

		r := ValueRange{lower: lower, upper: upper}

		ranges = append(ranges, r)
	}
	return ranges
}
func main() {
	sum := 0
	for _, r := range read_and_parse_input() {
		partial_sum := 0
		for i := r.lower; i <= r.upper; i++ {
			value_as_string := strconv.Itoa(i)
			if len(value_as_string)%2 != 0 {
				// Odd number of letters can not be two repetitions of the same number.
				continue
			}
			if value_as_string[:len(value_as_string)/2] == value_as_string[len(value_as_string)/2:] {
				partial_sum += i
			}
		}
		fmt.Println(partial_sum)
		sum += partial_sum
	}

	fmt.Println("Part 1", sum)
}
