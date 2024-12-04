package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type MatchFn func(i int, j int) bool

func main() {
	data, err := os.ReadFile("day04/example_2.txt")
	if err != nil {
		panic(err)
	}

	search := NewWordSearch(data)
	// part1 := search.Count(
	// 	search.Forward,
	// 	search.Upward,
	// 	search.DiagonallyBackward,
	// 	search.DiagonallyForward,
	// )
	// fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", search.Count(search.Cross))
}

type WordSearch struct {
	input [][]string
}

func NewWordSearch(input []byte) *WordSearch {
	parsed := make([][]string, 0)
	data := string(input)

	// Break lines
	lines := strings.Split(data, "\n")

	// Read backwards to make it easy to compare visually
	for _, line := range lines {
		current := make([]string, 0)
		for _, char := range line {
			current = append(current, string(char))
		}

		parsed = append(parsed, current)
	}

	return &WordSearch{input: parsed}
}

func (w *WordSearch) Forward(x int, y int) bool {
	// Out of boundaries
	if len(w.input[x]) < y+4 {
		return false
	}

	forward := strings.Join(w.input[x][y:y+4], "")

	return forward == "XMAS" || forward == "SAMX"
}

func (w *WordSearch) DiagonallyForward(x int, y int) bool {
	// Out of boundaries
	if len(w.input) < x+4 || len(w.input[x]) < y+4 {
		return false
	}

	df := ""

	for i := 0; i < 4; i++ {
		df += w.input[x+i][y+i]
	}

	return df == "XMAS" || df == "SAMX"
}

func (w *WordSearch) DiagonallyBackward(x int, y int) bool {
	// Out of boundaries
	if len(w.input) < x+4 || y-3 < 0 {
		return false
	}

	db := ""

	for i := 0; i < 4; i++ {
		db += w.input[x+i][y-i]
	}

	return db == "XMAS" || db == "SAMX"
}

func (w *WordSearch) Upward(x int, y int) bool {
	// Out of boundaries
	if x-3 < 0 {
		return false
	}

	upward := w.input[x][y] + w.input[x-1][y] + w.input[x-2][y] + w.input[x-3][y]

	return upward == "XMAS" || upward == "SAMX"
}

func (w *WordSearch) Cross(x int, y int) bool {
	// Skip if current character is not 'A'
	if w.input[x][y] != "A" {
		return false
	}

	// Out of boundaries
	if x-1 < 0 ||
		x+1 >= len(w.input) ||
		y+1 >= len(w.input[x]) ||
		y-1 < 0 {
		return false
	}

	topLeft := w.input[x-1][y-1]
	topRight := w.input[x-1][y+1]
	bottomLeft := w.input[x+1][y-1]
	bottomRight := w.input[x+1][y+1]

	result := (topLeft == "M" && bottomRight == "S" || topLeft == "S" && bottomLeft == "M") &&
		(topRight == "M" && bottomLeft == "S" || topRight == "S" || bottomRight == "M")

	if result {
		fmt.Println(topLeft, topRight, bottomLeft, bottomRight)
		fmt.Println("--------")
		fmt.Println(y, x)
		fmt.Println(topLeft, ".", topRight)
		fmt.Println(".", "A", ".")
		fmt.Println(bottomLeft, ".", bottomRight)
	}

	return result
}

func (w *WordSearch) Count(matchFns ...MatchFn) int {
	total := 0
	for i, lines := range w.input {
		for j := range lines {
			for _, fn := range matchFns {
				if fn(i, j) {
					total += 1
				}
			}
		}
	}

	return total
}

func (w *WordSearch) Print() {
	for _, lines := range slices.Backward(w.input) {
		fmt.Println(strings.Join(lines, ""))
	}
}
