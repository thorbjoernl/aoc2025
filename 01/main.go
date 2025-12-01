package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read_and_parse_input() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	instructions := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0] == 'R'
		value, _ := strconv.Atoi(line[1:])

		var instruction int
		if dir {
			instruction = value
		} else {
			instruction = -value
		}
		instructions = append(instructions, instruction)
	}
	return instructions
}

func main() {

	counter1 := 0
	dial := 50
	for _, delta := range read_and_parse_input() {
		dial += delta
		dial = dial % 100
		if dial == 0 {
			counter1++
		}
	}

	fmt.Println("Part 1:", counter1)
}
