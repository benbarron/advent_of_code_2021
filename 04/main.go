package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	value  int
	called bool
}

type Board struct {
	cells [][]Cell
}

func (b *Board) handleCall(value int) {
	for i := 0; i < len(b.cells); i++ {
		for j := 0; j < len(b.cells[0]); j++ {
			if b.cells[i][j].value == value {
				b.cells[i][j].called = true
			}
		}
	}
}

func (b *Board) isWinner() bool {
	if len(b.cells) == 0 {
		return false
	}

	var h int = len(b.cells)
	var w int = len(b.cells[0])

	for i := 0; i < h; i++ {
		var won bool = true
		for j := 0; j < w; j++ {
			if !b.cells[i][j].called {
				won = false
			}
		}
		if won {
			return true
		}
	}
	for i := 0; i < w; i++ {
		var won bool = true
		for j := 0; j < h; j++ {
			if !b.cells[j][i].called {
				won = false
			}
		}
		if won {
			return true
		}
	}
	return false
}

func (b *Board) calculateScore(lastCall int) int {
	var sum int = 0
	for i := 0; i < len(b.cells); i++ {
		for j := 0; j < len(b.cells[0]); j++ {
			if !b.cells[i][j].called {
				sum += b.cells[i][j].value
			}
		}
	}
	return sum * lastCall
}

func getInputData(filepath string) ([]int, []Board, error) {
	data, err := os.ReadFile(filepath)
	var calls []int
	var boards []Board

	if err != nil {
		err := fmt.Errorf("file %s does not exist", filepath)
		return calls, boards, err
	}

	var board_row_count int = 0
	calls = []int{}
	boards = []Board{}

	lines := strings.Split(string(data), "\n")

	// set game calls
	for _, num := range strings.Split(lines[0], ",") {
		v, _ := strconv.Atoi(num)
		calls = append(calls, v)
	}

	// read in boards
	var cells [][]Cell = [][]Cell{}

	for i := 1; i < len(lines); i++ {
		var inner []Cell = []Cell{}
		if lines[i] != "" {
			// continue adding to current board
			var board_row int = board_row_count % 5

			for _, num := range strings.Split(lines[i], " ") {
				if num != "" {
					v, _ := strconv.Atoi(num)
					cell := Cell{v, false}
					inner = append(inner, cell)
				}
			}
			cells = append(cells, inner)
			board_row += 1
		}
		if lines[i] == "" || i == len(lines)-1 {
			// end current board and start new one
			if len(cells) > 0 {
				var board Board = Board{cells}
				boards = append(boards, board)
			}
			cells = [][]Cell{}
		}
	}
	return calls, boards, nil
}

func partOne(calls []int, boards []Board) int {
	for _, call := range calls {
		for _, board := range boards {
			board.handleCall(call)
			if board.isWinner() {
				return board.calculateScore(call)
			}
		}
	}
	return -1
}

func partTwo(calls []int, boards []Board) int {
	var numBoards int = len(boards)
	var lastScore int = 0

	var winners []bool = make([]bool, numBoards)
	for i := 0; i < numBoards; i++ {
		winners[i] = false
	}

	for _, call := range calls {
		for i, board := range boards {
			if !winners[i] {
				board.handleCall(call)
				if board.isWinner() {
					lastScore = board.calculateScore(call)
					winners[i] = true
				}
			}
		}
	}
	return lastScore
}

func main() {

	if len(os.Args) < 2 {
		panic(fmt.Errorf("USAGE: go run ./04/main.go <input_file_path>"))
	}

	calls, boards, err := getInputData(os.Args[1])

	if err != nil {
		panic(err)
	}

	// fmt.Printf("Part 1: %d\n", partOne(calls, boards))
	fmt.Printf("Part 2: %d\n", partTwo(calls, boards))
}
