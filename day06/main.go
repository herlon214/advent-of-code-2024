package main

import (
	"fmt"
	"os"
	"strings"
)

type Position [2]int

type Tile string

const (
	Empty       Tile = ""
	Obstruction Tile = "#"
	Walked      Tile = "X"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type MovementResult int

const (
	Success MovementResult = iota
	Obstructed
	OutOfMap
)

func main() {
	data, err := os.ReadFile("day06/input.txt")
	if err != nil {
		panic(err)
	}

	m := NewMap(data)
	m.Walk()
	fmt.Println("Part 1:", len(m.steps))
}

type Map struct {
	tiles      [][]Tile
	currentPos Position
	direction  Direction
	steps      map[Position]bool
}

func NewMap(input []byte) *Map {
	lines := strings.Split(string(input), "\n")
	tiles := make([][]Tile, len(lines))
	startPos := Position{0, 0}

	for i, line := range lines {
		tiles[i] = make([]Tile, len(line))

		for j, tile := range line {
			if tile == '^' {
				startPos = Position{i, j}
				tiles[i][j] = Empty
			} else {
				tiles[i][j] = Tile(tile)
			}
		}
	}

	steps := make(map[Position]bool)
	steps[startPos] = true

	return &Map{
		tiles:      tiles,
		currentPos: startPos,
		steps:      steps,
	}
}

func (m *Map) Walk() {
	for {
		moveResult := m.MoveForward()

		switch moveResult {
		case OutOfMap:
			return
		case Obstructed:
			m.TurnRight()
		}
	}
}

func (m *Map) MoveForward() MovementResult {
	// Define next movement position
	nextPos := Position{m.currentPos[0], m.currentPos[1]}
	switch m.direction {
	case Up:
		nextPos = Position{m.currentPos[0] - 1, m.currentPos[1]}
	case Right:
		nextPos = Position{m.currentPos[0], m.currentPos[1] + 1}
	case Down:
		nextPos = Position{m.currentPos[0] + 1, m.currentPos[1]}
	case Left:
		nextPos = Position{m.currentPos[0], m.currentPos[1] - 1}
	}

	// Check boundaries
	if nextPos[0] < 0 ||
		nextPos[1] < 0 ||
		nextPos[0] >= len(m.tiles) ||
		nextPos[1] >= len(m.tiles[0]) {
		return OutOfMap
	}

	// Verify if next tile is an obstruction
	futureTile := m.tiles[nextPos[0]][nextPos[1]]
	if futureTile == Obstruction {
		return Obstructed
	}

	// Update position
	m.currentPos = nextPos

	// Save step
	m.steps[m.currentPos] = true

	return Success
}

func (m *Map) CurrentTile() Tile {
	return m.tiles[m.currentPos[0]][m.currentPos[1]]
}

func (m *Map) TurnRight() {
	switch m.direction {
	case Up:
		m.direction = Right
	case Right:
		m.direction = Down
	case Down:
		m.direction = Left
	case Left:
		m.direction = Up
	}
}

func (m *Map) Print() {
	for i := range m.tiles {
		for j := range m.tiles[i] {
			// Print current position
			if m.currentPos[0] == i && m.currentPos[1] == j {
				fmt.Print(m.direction.String())
			} else if m.steps[Position{i, j}] {
				fmt.Print(Walked)
			} else {
				fmt.Print(m.tiles[i][j])
			}
		}
		fmt.Print("\n")
	}
}

func (d Direction) String() string {
	switch d {
	case Up:
		return "^"
	case Right:
		return ">"
	case Down:
		return "v"
	case Left:
		return "<"
	default:
		panic("unknown direction")
	}
}
