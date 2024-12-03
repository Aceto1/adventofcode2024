package main

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("./01.txt")

	check(err)

	inputString := string(dat)

	// Normalize line endings
	inputString = strings.ReplaceAll(inputString, "\r\n", "\n")

	lines := strings.Split(inputString, "\n")
	leftList := []int{}
	rightList := []int{}

	for i := range lines {
		lineParts := strings.Split(strings.TrimSpace(lines[i]), " ")

		leftInt, err := strconv.Atoi(lineParts[0])
		check(err)
		leftList = append(leftList, leftInt)

		rightInt, err := strconv.Atoi(lineParts[len(lineParts)-1])
		check(err)
		rightList = append(rightList, rightInt)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDiff := 0

	// lists are of same length so no need to zip the lists
	for i := range leftList {
		// numbers are always integers, so no precision is lost by casting to and from float64
		totalDiff += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	print("The total difference of the two lists is: " + strconv.Itoa(totalDiff))
	os.Exit(0)
}
