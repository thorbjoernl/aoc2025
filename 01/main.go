package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type rotation struct {
	direction bool // R: True, L: False
	value     int
}

func read_and_parse_input() []rotation {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	instructions := make([]rotation, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0] == 'R'
		value, _ := strconv.Atoi(line[1:])

		instructions = append(instructions, rotation{direction: dir, value: value})
	}
	return instructions
}

func main() {

	counter := 0
	dial := 50
	for _, rot := range read_and_parse_input() {
		if rot.direction {
			dial += rot.value
		} else {
			dial -= rot.value
		}
		dial = dial % 100
		if dial == 0 {
			counter++
		}
	}

	fmt.Println(counter)
}
