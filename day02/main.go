package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const MaxDiff = 3

func main() {
	data, err := os.ReadFile("day02/input.txt")
	if err != nil {
		panic(err)
	}

	reports := CreateReports(data)

	fmt.Println("Part 1:", reports.StrictlySafeCount())
	fmt.Println("Part 2:", reports.FlexiblySafeCount())

}

type Reports []Report

func CreateReports(data []byte) Reports {
	lines := strings.Split(string(data), "\n")
	reports := make([]Report, 0)

	for _, line := range lines {
		lineNumbers := strings.Split(line, " ")

		nLine := make([]int, 0)

		for _, number := range lineNumbers {
			number, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}

			nLine = append(nLine, number)
		}

		reports = append(reports, Report{levels: nLine})
	}

	return reports
}

func (r Reports) StrictlySafeCount() int {
	safe := 0
	for _, report := range r {
		if report.StrictlySafe() {
			safe++
		}
	}

	return safe
}

func (r Reports) FlexiblySafeCount() int {
	safe := 0
	for _, report := range r {
		if report.FlexiblySafe() {
			safe++
		}
	}

	return safe
}

type Report struct {
	levels []int
}

func (r Report) FlexiblySafe() bool {
	// Generate new levels removing one possible outlier
	levels := make([][]int, 0)
	levels = append(levels, r.levels)

	for i := range r.levels {
		newLevel := make([]int, 0)

		newLevel = append(newLevel, r.levels[:i]...)
		if i < len(r.levels) {
			newLevel = append(newLevel, r.levels[i+1:]...)
		}

		levels = append(levels, newLevel)
	}

	// Verify all levels
	for _, level := range levels {
		report := Report{levels: level}
		if report.StrictlySafe() {
			return true
		}
	}

	return false
}

func (r Report) StrictlySafe() bool {
	var bit *bool

	for i := 0; i < len(r.levels)-1; i++ {
		if r.levels[i] == r.levels[i+1] {
			return false
		}

		diff := float64(r.levels[i]) - float64(r.levels[i+1])
		sign := math.Signbit(diff)

		if math.Abs(diff) > MaxDiff {
			return false
		}

		if bit == nil {
			bit = &sign
		} else {
			if *bit != sign {
				return false
			}
		}

	}

	return true
}
