package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Tile string

const (
	Empty     Tile = "."
	Antinode  Tile = "#"
	Highlight Tile = "@"
)

type Position [2]int

type Map [][]Tile

type Scanner struct {
	Map           Map
	antennas      map[Tile][]Position
	totalAntennas int
	antinodes     map[Position]bool
}

func main() {
	data, err := os.ReadFile("day08/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := NewScanner(data)
	scanner.FindAntinodes(1, false)
	fmt.Println("Part 1:", scanner.CountAntinodes())

	scanner = NewScanner(data)
	scanner.FindAntinodes(math.MaxInt, true)
	fmt.Println("Part 2:", scanner.CountAntinodes())
}

func NewScanner(input []byte) *Scanner {
	m := make(Map, 0)
	antennas := make(map[Tile][]Position, 0)
	totalAntennas := 0

	for i, line := range strings.Split(string(input), "\n") {
		mm := make([]Tile, 0)

		for j, char := range line {
			freq := Tile(char)

			// Map antennas
			if freq != Empty {
				if _, ok := antennas[freq]; !ok {
					antennas[freq] = make([]Position, 0)
				}
				antennas[freq] = append(antennas[freq], Position{i, j})

				// Increase total number of antennas
				totalAntennas += 1
			}

			mm = append(mm, freq)
		}

		m = append(m, mm)
	}

	return &Scanner{
		Map:           m,
		antennas:      antennas,
		antinodes:     make(map[Position]bool),
		totalAntennas: totalAntennas,
	}
}

// FindAtinodes on each direction specifying
// a max number of antinodes for both directions
func (s *Scanner) FindAntinodes(maxAntinodes int, includeAntennas bool) {
	// Check all antennas with same frequency
	for _, antennas := range s.antennas {
		visited := make(map[Position]bool)

		// Combine them
		for i := 0; i < len(antennas); i++ {
			for j := 0; j < len(antennas); j++ {
				// Skip itself and visited positions
				if i == j || visited[Position{i, j}] {
					continue
				}

				lhs := antennas[i]
				rhs := antennas[j]

				diffA := lhs[0] - rhs[0]
				diffB := lhs[1] - rhs[1]
				diffPos := Position{diffA, diffB}

				lhsAntinode := lhs.Add(diffPos)
				rhsAntinode := rhs.Sub(diffPos)

				added := 0
				for {
					lhsOOB := s.Map.IsPositionOutOfBounds(lhsAntinode)
					rhsOOB := s.Map.IsPositionOutOfBounds(rhsAntinode)

					if (lhsOOB && rhsOOB) || added == maxAntinodes {
						break
					}

					// Add antinodes to map if it doesn't overlap
					if !lhsOOB {
						s.antinodes[lhsAntinode] = true

						// Bump current position
						lhsAntinode = lhsAntinode.Add(diffPos)
					}
					if !rhsOOB {
						s.antinodes[rhsAntinode] = true

						// Bump current position
						rhsAntinode = rhsAntinode.Sub(diffPos)
					}

					added += 1
				}

				// If added at least one then mark current nodes as antinodes
				if includeAntennas {
					s.antinodes[lhs] = true
					s.antinodes[rhs] = true
				}

				// Mark as visited from both antennas' perspective
				visited[Position{i, j}] = true
				visited[Position{j, i}] = true

			}
		}
	}
}

func (s *Scanner) PrintMap(highlight ...Position) {
	for i, tiles := range s.Map {
		for j, tile := range tiles {
			// Check antinode
			_, hasAntinode := s.antinodes[Position{i, j}]

			// Print antinode
			if hasAntinode {
				fmt.Printf("%s", Antinode)
			} else if slices.Contains(highlight, Position{i, j}) { // Print hightlight
				fmt.Printf("%s", Highlight)
			} else { // Print tile
				fmt.Printf("%s", tile)
			}
		}
		fmt.Printf("\n")
	}
}

func (s *Scanner) CountAntinodes() int {
	return len(s.antinodes)
}

func (p Position) Sub(pos Position) Position {
	return Position{p[0] - pos[0], p[1] - pos[1]}
}

func (p Position) Add(pos Position) Position {
	return Position{p[0] + pos[0], p[1] + pos[1]}
}

func (m Map) IsPositionOutOfBounds(pos Position) bool {
	return pos[0] < 0 ||
		pos[0] >= len(m) ||
		pos[1] < 0 ||
		pos[1] >= len(m[0])
}
