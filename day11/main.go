package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day11/input.txt")
	if err != nil {
		panic(err)
	}

	valuesStr := strings.Split(string(data), " ")
	values := make([]int, 0)

	for _, str := range valuesStr {
		value, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		values = append(values, value)
	}

	fmt.Println("Part 1:", Blink(values, 25))
}

func Blink(values []int, times int) int {
	currentValues := slices.Clone(values)

	for i := 0; i < times; i++ {
		newValues := make([]int, 0)

		for _, val := range currentValues {
			// 1st rule
			if val == 0 {
				newValues = append(newValues, 1)

				continue
			}

			// 2nd rule
			digits := fmt.Sprint(val)
			if len(digits)%2 == 0 {
				half := len(digits) / 2
				leftHalf, err := strconv.Atoi(digits[:half])
				if err != nil {
					panic(err)
				}
				rightHalf, err := strconv.Atoi(digits[half:])
				if err != nil {
					panic(err)
				}

				newValues = append(newValues, leftHalf, rightHalf)

				continue
			}

			newValues = append(newValues, val*2024)
		}

		currentValues = newValues

	}

	return len(currentValues)
}
