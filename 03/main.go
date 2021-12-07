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

func filterNthBit(gauge string, bitNumber int, data [][]int) [][]int {
	if len(data) == 1 {
		return data
	}

	var count [2]int = [2]int{0, 0}
	var filteredData [][]int = [][]int{}

	for j := 0; j < len(data); j++ {
		count[data[j][bitNumber]]++
	}

	if gauge == "o2" {
		if count[0] > count[1] {
			for i := 0; i < len(data); i++ {
				if data[i][bitNumber] == 0 {
					filteredData = append(filteredData, data[i])
				}
			}
		} else {
			for i := 0; i < len(data); i++ {
				if data[i][bitNumber] == 1 {
					filteredData = append(filteredData, data[i])
				}
			}
		}
	} else {
		if count[0] > count[1] {
			for i := 0; i < len(data); i++ {
				if data[i][bitNumber] == 1 {
					filteredData = append(filteredData, data[i])
				}
			}
		} else {
			for i := 0; i < len(data); i++ {
				if data[i][bitNumber] == 0 {
					filteredData = append(filteredData, data[i])
				}
			}
		}
	}

	return filterNthBit(gauge, bitNumber+1, filteredData)
}

func partTwo(data [][]int) int {
	oxygenData := filterNthBit("o2", 0, data)
	carbonData := filterNthBit("co2", 0, data)

	var oxygenBinary string = ""
	var carbonBinary string = ""

	for _, e := range oxygenData[0] {
		oxygenBinary += fmt.Sprintf("%d", e)
	}
	for _, e := range carbonData[0] {
		carbonBinary += fmt.Sprintf("%d", e)
	}

	oxygenRating, _ := strconv.ParseInt(oxygenBinary, 2, 64)
	carbonRating, _ := strconv.ParseInt(carbonBinary, 2, 64)

	return int(carbonRating) * int(oxygenRating)
}

func main() {

	if len(os.Args) < 2 {
		panic(fmt.Errorf("USAGE: go run ./03/main.go <input_file_path>"))
	}

	data, err := getInputData(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", partOne(data))
	fmt.Printf("Part 2: %d\n", partTwo(data))
}
