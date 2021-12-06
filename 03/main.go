package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputData(filepath string) ([][]int, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		err := fmt.Errorf("file %s does not exist", filepath)
		var data [][]int
		return data, err
	}
	var lines []string = strings.Split(string(data), "\n")
	var cols int = len(lines[0])
	var rows int = len(lines)

	// fmt.Printf("%d, %d\n", cols, rows)

	var array [][]int
	for i := 0; i < rows; i++ {
		var inner []int = []int{}
		var line []string = strings.Split(lines[i], "")
		for j := 0; j < cols; j++ {
			v, _ := strconv.Atoi(line[j])
			inner = append(inner, v)
		}
		array = append(array, inner)
	}
	return array, nil
}

func partOne(data [][]int) int {
	var gammaRateBinary string = ""
	var epsilonRateBinary string = ""
	for i := 0; i < len(data[0]); i++ {
		var count [2]int = [2]int{0, 0}
		for j := 0; j < len(data); j++ {
			count[data[j][i]]++
		}
		if count[0] > count[1] {
			gammaRateBinary += "0"
			epsilonRateBinary += "1"
		} else {
			gammaRateBinary += "1"
			epsilonRateBinary += "0"
		}
	}
	gammaRate, _ := strconv.ParseInt(gammaRateBinary, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonRateBinary, 2, 64)
	return int(gammaRate) * int(epsilonRate)
}

func partTwo() int {

	return 0
}

func main() {

	if len(os.Args) < 2 {
		panic(fmt.Errorf("USAGE: go run ./03/main.go <input_file_path>"))
	}

	data, err := getInputData(os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, v := range data {
		fmt.Println(v)
	}
	fmt.Println("-----------------------------------------------------")

	fmt.Printf("Part 1: %d\n", partOne(data))
	// fmt.Printf("Part 2: %d\n", partTwo(commands))
}
