package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ferror(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dat, err := os.ReadFile("input.txt")
	ferror(err)
	lines := strings.Split(string(dat), "\n")
	part1, part2 := 0, 0

	for idx, l := range lines {
		g := strings.Split(l, ":")
		if len(g) < 2 {
			continue
		}
		if possible(g[1]) {
			part1 += idx + 1
		}

		part2 += power(g[1])
	}

	fmt.Printf("Part 1: %d, Part 2: %d", part1, part2)
}

func possible(game string) bool {
	max := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	rounds := strings.Split(game, ";")
	for _, r := range rounds {
		colors := strings.Split(r, ",")
		for _, c := range colors {
			c := strings.Trim(c, " ")
			count := strings.Split(c, " ")
			iCount, _ := strconv.Atoi(count[0])
			if max[count[1]] < iCount {
				return false
			}

		}
	}

	return true
}

func power(game string) int {
	rounds := strings.Split(game, ";")
	fCount := make(map[string]int)

	for _, r := range rounds {
		colors := strings.Split(r, ",")
		for _, c := range colors {
			c := strings.Trim(c, " ")
			count := strings.Split(c, " ")
			iCount, _ := strconv.Atoi(count[0])
			if fCount[count[1]] < iCount {
				fCount[count[1]] = iCount
			}

		}

	}
	product := 1
	for _, f := range fCount {
		product *= f

	}

	return product
}
