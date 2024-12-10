package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Position [2]int

func (p Position) Add(n Position) Position {
	return Position{
		p[0] + n[0],
		p[1] + n[1],
	}
}

type Map struct {
	tiles      [][]int
	trailHeads []Position
}

func main() {
	data, err := os.ReadFile("day10/input.txt")
	if err != nil {
		panic(err)
	}

	hikeMap := NewMap(data)
	fmt.Println("Part 1:", hikeMap.Walk(true))
	fmt.Println("Part 2:", hikeMap.Walk(false))
}

func NewMap(data []byte) *Map {
	m := make([][]int, 0)
	heads := make([]Position, 0)

	for i, line := range strings.Split(string(data), "\n") {
		tiles := make([]int, 0)
		for j, tile := range line {
			// Empty
			if tile == '.' {
				tiles = append(tiles, -1)

				continue
			}

			val, err := strconv.Atoi(string(tile))
			if err != nil {
				panic(err)
			}

			if val == 0 {
				heads = append(heads, Position{i, j})
			}

			tiles = append(tiles, val)
		}

		m = append(m, tiles)
	}

	return &Map{
		tiles:      m,
		trailHeads: heads,
	}
}

func (m *Map) Walk(distinct bool) int {
	scores := 0

	for _, head := range m.trailHeads {
		values := make([]int, 0)
		steps := make([]Position, 0)
		steps = append(steps, head)

		for {
			nextSteps := make([]Position, 0)

			for _, step := range steps {
				possibleSteps := m.PossibleSteps(step)

				// Filter existing
				for _, possibleStep := range possibleSteps {
					if distinct {
						if !slices.Contains(nextSteps, possibleStep) {
							nextSteps = append(nextSteps, possibleStep)
						}
					} else {
						nextSteps = append(nextSteps, possibleStep)
					}
				}
			}

			if len(nextSteps) == 0 {
				break
			}

			steps = nextSteps
		}

		// Calculate how far we reached
		for _, step := range steps {
			val := m.Value(step)
			if val == 9 {
				values = append(values, m.Value(step))
			}
		}

		scores += len(values)
	}

	return scores
}

func (m *Map) PossibleSteps(from Position) []Position {
	possible := make([]Position, 0)
	currentVal := m.Value(from)

	// Up
	up := from.Add(Position{-1, 0})
	if !m.IsOutOfBounds(up) && m.Value(up) == currentVal+1 {
		possible = append(possible, up)
	}

	// Down
	down := from.Add(Position{1, 0})
	if !m.IsOutOfBounds(down) && m.Value(down) == currentVal+1 {
		possible = append(possible, down)
	}

	// Right
	right := from.Add(Position{0, 1})
	if !m.IsOutOfBounds(right) && m.Value(right) == currentVal+1 {
		possible = append(possible, right)
	}

	// Left
	left := from.Add(Position{0, -1})
	if !m.IsOutOfBounds(left) && m.Value(left) == currentVal+1 {
		possible = append(possible, left)
	}

	return possible
}

func (m *Map) Value(pos Position) int {
	return m.tiles[pos[0]][pos[1]]
}

func (m *Map) IsOutOfBounds(pos Position) bool {
	return pos[0] < 0 || pos[1] < 0 || pos[0] >= len(m.tiles) || pos[1] >= len(m.tiles[0])
}

func (m *Map) Print() {
	for _, tiles := range m.tiles {
		for _, tile := range tiles {

			if tile == -1 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", tile)
			}
		}

		fmt.Printf("\n")

	}
}
