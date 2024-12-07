package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Operator string

const (
	Multiply    Operator = "*"
	Add         Operator = "+"
	Concatenate Operator = "||"
)

type Equation struct {
	result             int64
	numbers            []int64
	possibleOperations [][]Operator
}

type Equations []Equation

func main() {
	data, err := os.ReadFile("day07/input.txt")
	if err != nil {
		panic(err)
	}

	// Part 1
	eqPart1 := ParseEquations(data, []Operator{Add, Multiply})
	fmt.Println("Part 1:", eqPart1.FilterPossible().TotalCalibrationResult())

	// Part 2
	eqPart2 := ParseEquations(data, []Operator{Add, Multiply, Concatenate})
	fmt.Println("Part 2:", eqPart2.FilterPossible().TotalCalibrationResult())
}

func ParseEquations(input []byte, availableOperators []Operator) Equations {
	equations := make(Equations, 0)

	for _, line := range strings.Split(string(input), "\n") {
		// Test result
		result, err := strconv.Atoi(strings.Split(line, ":")[0])
		if err != nil {
			panic(err)
		}

		// Numbers to be used
		numbersStr := strings.Split(strings.Split(line, ": ")[1], " ")
		numbers := make([]int64, 0)

		for _, numberStr := range numbersStr {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				panic(err)
			}

			numbers = append(numbers, int64(number))
		}

		// Generate all possible combinations of +, * and || operators
		numOperators := len(numbers) - 1
		possibleOperations := make([][]Operator, 0)

		// Helper function to generate permutations
		var generateOperations func(current []Operator)
		generateOperations = func(current []Operator) {
			if len(current) == numOperators {
				possibleOperations = append(possibleOperations, slices.Clone(current))
				return
			}

			// Try each operator
			for _, op := range availableOperators {
				current = append(current, op)
				generateOperations(current)
				current = current[:len(current)-1]
			}
		}

		generateOperations(make([]Operator, 0))

		equations = append(equations, Equation{
			result:             int64(result),
			numbers:            numbers,
			possibleOperations: possibleOperations,
		})
	}

	return equations
}

func (e Equations) FilterPossible() Equations {
	result := make(Equations, 0)

	for _, eq := range e {
		if eq.IsPossible() {
			result = append(result, eq)
		}
	}

	return result
}

func (e Equations) TotalCalibrationResult() int64 {
	total := int64(0)
	for _, eq := range e {
		total += eq.result
	}

	return total
}

func (e Equation) IsPossible() bool {
	for _, operators := range e.possibleOperations {
		if e.Evaluate(operators) == e.result {
			return true
		}
	}

	return false
}

func (e Equation) Evaluate(operators []Operator) int64 {
	// Resolve left to right
	numbers := slices.Clone(e.numbers)
	left := numbers[0]
	ops := slices.Clone(operators)

	for {
		if len(numbers) == 0 || len(ops) == 0 {
			break
		}
		right := numbers[1]
		op := ops[0]
		result := op.Execute(left, right)

		left = result

		// Bump sequence
		ops = ops[1:]
		numbers = numbers[1:]
	}

	return left
}

func (o Operator) Execute(a int64, b int64) int64 {
	switch o {
	case Multiply:
		return a * b
	case Add:
		return a + b
	case Concatenate:
		val, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
		if err != nil {
			panic(err)
		}

		return int64(val)
	default:
		panic("unknown operator")
	}
}
