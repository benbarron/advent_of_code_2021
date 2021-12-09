package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BOARD_SIZE int = 1000

type Point struct {
	x int
	y int
}

type Vector struct {
	start Point
	end   Point
}

func getInputData(filepath string) ([]Vector, error) {
	data, err := os.ReadFile(filepath)
	var vectors []Vector = make([]Vector, 0)

	if err != nil {
		err := fmt.Errorf("file %s does not exist", filepath)
		return vectors, err
	}

	for _, line := range strings.Split(string(data), "\n") {
		split_string := strings.Split(line, "->")
		p1_strings := strings.Split(split_string[0], ",")
		p2_strings := strings.Split(split_string[1], ",")
		x1, _ := strconv.Atoi(strings.TrimSpace(p1_strings[0]))
		y1, _ := strconv.Atoi(strings.TrimSpace(p1_strings[1]))
		x2, _ := strconv.Atoi(strings.TrimSpace(p2_strings[0]))
		y2, _ := strconv.Atoi(strings.TrimSpace(p2_strings[1]))
		var point1 Point = Point{x1, y1}
		var point2 Point = Point{x2, y2}
		var vector Vector = Vector{point1, point2}
		vectors = append(vectors, vector)
	}
	return vectors, nil
}

func allocBoard() [][]int {
	var board [][]int = [][]int{}
	for i := 0; i < BOARD_SIZE; i++ {
		var inner []int = []int{}
		for j := 0; j < BOARD_SIZE; j++ {
			inner = append(inner, 0)
		}
		board = append(board, inner)
	}
	return board
}

func printBoard(board [][]int) {
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if board[i][j] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", board[i][j])
			}
		}
		fmt.Printf("\n")
	}
}

func printVectors(vectors []Vector) {
	for _, vector := range vectors {
		fmt.Printf("(%d, %d) -> ", vector.start.x, vector.start.y)
		fmt.Printf("(%d, %d)\n", vector.end.x, vector.end.y)
	}
}

func partOne(vectors []Vector) int {
	board := allocBoard()
	for _, line := range vectors {
		if line.start.x == line.end.x {
			if line.start.y < line.end.y {
				for i := line.start.y; i <= line.end.y; i++ {
					board[i][line.start.x]++
				}
			}

			if line.start.y > line.end.y {
				for i := line.end.y; i <= line.start.y; i++ {
					board[i][line.start.x]++
				}
			}

			if line.start.y == line.end.y {
				board[line.start.x][line.end.x]++
			}
		} else if line.start.y == line.end.y {
			if line.start.x < line.end.x {
				for i := line.start.x; i <= line.end.x; i++ {
					board[line.start.y][i]++
				}
			}

			if line.start.x > line.end.x {
				for i := line.end.x; i <= line.start.x; i++ {
					board[line.start.y][i]++
				}
			}
		}
	}

	var count int = 0
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if board[i][j] >= 2 {
				count++
			}
		}
	}
	return count
}

func partTwo(data []Vector) int {
	return 0
}

func main() {

	if len(os.Args) < 2 {
		panic(fmt.Errorf("USAGE: go run ./05/main.go <input_file_path>"))
	}

	data, err := getInputData(os.Args[1])

	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", partOne(data))
	fmt.Printf("Part 2: %d\n", partTwo(data))
}
