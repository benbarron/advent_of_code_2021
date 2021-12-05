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

//199  A			1
//200  A B			1,2
//208  A B C		1,2,3
//210    B C D		2,3,4
//200  E   C D		1,3,4
//207  E F   D
//240  E F G
//269    F G H
//260      G H
//263        H


func part_one(data []string) int {
	prev := 0
	inc_count := 0

	for i, v := range data {
		measurement, _ := strconv.Atoi(v)
		if measurement > prev && i > 0 {
			inc_count++
		}
		prev = measurement
	}
	return inc_count
}

func part_two(data []string) int {
	prev := 0
	inc_count := 0

	for i:= 0; i < len(data) - 2 ; i++ {
		v1, _ := strconv.Atoi(data[i])
		v2, _ := strconv.Atoi(data[i+1])
		v3, _ := strconv.Atoi(data[i+2])
		measurement := v1+v2+v3

		if measurement > prev && i > 0 {
			inc_count++
		}
		prev = measurement
	}
	return inc_count
}

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("USAGE: go run ./1_sonar_sweep.go <input_file_path>"))
	}

	data, err := getInputData(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part_one(data))
	fmt.Printf("Part 1: %d\n", part_two(data))
}
