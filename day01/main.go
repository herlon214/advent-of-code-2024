package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	example, err := os.ReadFile("day01/example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Example result:", part1(example))

	input, err := os.ReadFile("day01/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func part1(data []byte) int {
	lines := strings.Split(string(data), "\n")
	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range lines {
		numbers := strings.Split(line, " ")

		number, _ := strconv.Atoi(numbers[0])
		left = append(left, number)

		number, _ = strconv.Atoi(numbers[len(numbers)-1])
		right = append(right, number)
	}

	// Sort both lists
	sort.Ints(left)
	sort.Ints(right)

	// Find the distance between them
	total := 0
	for i := 0; i < len(left); i++ {
		diff := math.Abs(float64(left[i]) - float64(right[i]))
		total += int(diff)
	}

	return total
}

func part2(data []byte) int {
	lines := strings.Split(string(data), "\n")
	left := make([]int, 0)
	right := make(map[int]int, 0)

	for _, line := range lines {
		numbers := strings.Split(line, " ")

		number, _ := strconv.Atoi(numbers[0])
		left = append(left, number)

		number, _ = strconv.Atoi(numbers[len(numbers)-1])
		right[number] += 1
	}

	total := 0

	for _, num := range left {
		total += num * right[num]
	}

	return total
}
