package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("day03/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func part2(data []byte) int {
	// Append an initial do()
	input := `do()` + string(data)

	mulReg := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	permissionReg := regexp.MustCompile(`(do|don't)\(\)`)

	permissions := permissionReg.FindAllStringSubmatchIndex(input, -1)
	operations := mulReg.FindAllStringSubmatchIndex(input, -1)

	total := 0
	for {
		if len(operations) == 0 {
			break
		}

		left, err := strconv.Atoi(input[operations[0][2]:operations[0][3]])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(input[operations[0][4]:operations[0][5]])
		if err != nil {
			panic(err)
		}

		// Bump permission if needed
		if len(permissions) > 1 {
			nextPermission := permissions[1]

			if operations[0][0] > nextPermission[0] {
				permissions = permissions[1:]
			}
		}

		// Check current permission
		permission := input[permissions[0][0]:permissions[0][1]]
		if permission == `do()` {
			total += left * right
		}

		operations = operations[1:]
	}

	return total
}

func part1(data []byte) int {
	mulReg := regexp.MustCompile("mul\\(([0-9]+),([0-9]+)\\)")

	operations := mulReg.FindAllStringSubmatch(string(data), -1)

	total := 0
	for _, match := range operations {
		left, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		right, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		total += left * right
	}

	return total
}
