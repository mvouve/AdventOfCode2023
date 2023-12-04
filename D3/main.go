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
	p2(lines)

}

func p2(lines []string) {
	stars := make(map[string]*[]int)
	sum := 0

	// start by finding all the stars.
	for row, l := range lines {
		for column, c := range l {
			if c == '*' {
				fmt.Println(row, column)
				arr := make([]int, 0, 2)
				stars[fmt.Sprintf("%d,%d", row, column)] = &arr
			}
		}
	}

	// find all the numbers
	findAllNumbersAndCallback(lines, func(start, end, row int) {
		// see if there's * touching it.
		for y := row - 1; y <= row+1; y++ {
			// end will already be + 1.
			for x := start - 1; x <= end; x++ {
				if stars[fmt.Sprintf("%d,%d", y, x)] != nil {
					num, _ := strconv.Atoi(string(lines[row][start:end]))
					n := append(*stars[fmt.Sprintf("%d,%d", y, x)], num)
					stars[fmt.Sprintf("%d,%d", y, x)] = &n
				}
			}
		}
	})
	// finally, multiply the gears together
	for c, s := range stars {
		fmt.Println(c, s)
		if len(*s) == 2 {
			a := *s
			sum += (a[0] * a[1])
		}
	}

	fmt.Printf("Sum = %d", sum)
}

func p1(lines []string) {
	sum := 0
	findAllNumbersAndCallback(lines, func(start, c, row int) {
		if symbolSeek(lines, start, c, row) {
			num, _ := strconv.Atoi(string(lines[row][start:c]))
			sum += num
		}
	})

	fmt.Printf("Sum = %d", sum)

}

func findAllNumbersAndCallback(lines []string, callback func(start, c, row int)) {
	for row, l := range lines {
		for c := 0; c < len(l); c += 1 {
			start := c
			for c < len(l) && (l[c] >= '0' && l[c] <= '9') {
				c++
			}
			if c != start {
				callback(start, c, row)
			}
		}
	}

}

func symbolSeek(lines []string, startColumn, stopColumn, row int) bool {

	x := startColumn - 1
	if x < 0 {
		x = 0
	}
	maxX := stopColumn
	if maxX >= len(lines[row])-1 {
		maxX = len(lines[row]) - 1
	}
	y := row - 1
	if y < 0 {
		y = 0
	}
	maxY := row + 1
	if maxY >= len(lines)-1 {

		maxY = row
	}

	minX := x
	fmt.Printf("\n\n")
	for y <= maxY {
		for x <= maxX {
			fmt.Printf("%c", lines[y][x])
			// We'll just check that it's not a number or a .
			if (lines[y][x] < '0' || lines[y][x] > '9') && lines[y][x] != '.' {
				return true
			}
			x++
		}
		fmt.Printf("\n")
		y++
		x = minX
	}

	return false
}
