package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputData(filepath string) ([]string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		err := fmt.Errorf("file %s does not exist", filepath)
		lines := []string{}
		return lines, err
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("USAGE: go run ./1_sonar_sweep.go <input_file_path>"))
	}

	data, err := getInputData(os.Args[1])
	if err != nil {
		panic(err)
	}

	prev := 0
	inc_count := 0

	for i, v := range data {
		measurement, _ := strconv.Atoi(v)
		if measurement > prev && i > 0 {
			inc_count++
		}
		prev = measurement
	}

	fmt.Printf("Number of measure increases: %d\n", inc_count)
}
