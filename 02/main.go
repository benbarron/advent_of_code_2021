package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	direction string
	value     int
}

func getInputData(filepath string) ([]Command, error) {
	data, err := os.ReadFile(filepath)
	var commands []Command
	if err != nil {
		err := fmt.Errorf("file %s does not exist", filepath)
		return commands, err
	}
	commands = []Command{}
	for _, v := range strings.Split(string(data), "\n") {
		var line []string = strings.Split(v, " ")
		value, _ := strconv.Atoi(line[1])
		newCommand := Command{line[0], value}
		commands = append(commands, newCommand)
	}
	return commands, nil
}

func partOne(commands []Command) int {
	var x int = 0
	var y int = 0

	for _, command := range commands {
		switch command.direction {
		case "forward":
			x += command.value
		case "up":
			y -= command.value
		case "down":
			y += command.value
		}
	}

	return x * y
}

func partTwo(commands []Command) int {
	var x int = 0
	var y int = 0
	var aim int = 0

	for _, command := range commands {
		switch command.direction {
		case "forward":
			x += command.value
			y += (aim * command.value)
		case "up":
			aim -= command.value
		case "down":
			aim += command.value
		}
		fmt.Printf("%s\t-> x=%d, y=%d, a=%d\n", command.direction, x, y, aim)
	}

	return x * y
}

func main() {

	if len(os.Args) < 2 {
		panic(fmt.Errorf("USAGE: go run ./02/main.go <input_file_path>"))
	}

	commands, err := getInputData(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", partOne(commands))
	fmt.Printf("Part 2: %d\n", partTwo(commands))
}
