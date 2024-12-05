package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PrintQueue struct {
	before  map[int]map[int]bool
	after   map[int]map[int]bool
	updates [][]int
}

func main() {
	data, err := os.ReadFile("day05/input.txt")
	if err != nil {
		panic(err)
	}

	pq := NewPrintQueue(string(data))
	fmt.Println("Part 1:", pq.SumCorrect())
	fmt.Println("Part 2:", pq.SumWrong())
}

func NewPrintQueue(input string) *PrintQueue {
	// Parse lines
	lines := strings.Split(string(input), "\n")

	before := make(map[int]map[int]bool)
	after := make(map[int]map[int]bool)
	updates := make([][]int, 0)

	for _, line := range lines {
		// Update graph with ordering rules
		if strings.Contains(line, "|") {
			numbers := strings.Split(line, "|")

			// Convert numbers
			left, err := strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}

			right, err := strconv.Atoi(numbers[1])
			if err != nil {
				panic(err)
			}

			// Check if graph item is created
			if _, ok := before[left]; !ok {
				before[left] = make(map[int]bool, 0)
			}
			if _, ok := after[right]; !ok {
				after[right] = make(map[int]bool, 0)
			}

			before[left][right] = true
			after[right][left] = true
		}

		// Insert new update
		if strings.Contains(line, ",") {
			numbers := make([]int, 0)
			items := strings.Split(line, ",")

			for _, item := range items {
				number, err := strconv.Atoi(item)
				if err != nil {
					panic(err)
				}

				numbers = append(numbers, number)
			}

			updates = append(updates, numbers)

		}
	}

	return &PrintQueue{
		before:  before,
		after:   after,
		updates: updates,
	}

}

func (p *PrintQueue) SumCorrect() int {
	total := 0

	for _, items := range p.updates {
		wrong := false

		for i := 0; i < len(items)-1; i++ {
			left := items[i]
			right := items[i+1]

			// Check if it's in the wrong order
			if _, ok := p.before[left][right]; !ok {
				wrong = true
				break
			}
		}

		if !wrong {
			middle := items[len(items)/2]
			total += middle
		}
	}

	return total
}

func (p *PrintQueue) SumWrong() int {
	total := 0
	retries := make(map[int]int)

	i := 0

	for {
		if i == len(p.updates) {
			break
		}

		items := p.updates[i]

		j := 0
		swapCount := 0
		for {
			if j == len(items)-1 {
				break
			}

			left := items[j]
			right := items[j+1]

			if _, ok := p.before[left][right]; !ok {
				// Check if swap is allowed
				if p.after[left][right] && p.before[right][left] {
					// Swap values
					items[j], items[j+1] = items[j+1], items[j]

					// Do not update j try and try again
					retries[i]++
					swapCount++
					continue
				}

			}

			j++
		}

		// A swap happened so we need to verify the current items again
		if swapCount > 0 {
			continue
		}

		// No swap anymore but it was fixed
		if swapCount == 0 && retries[i] > 0 {
			middle := items[len(items)/2]
			total += middle
		}

		i += 1
	}

	return total
}
