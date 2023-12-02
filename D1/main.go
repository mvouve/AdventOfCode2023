package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

)

func sub(s string) string {
	for o, r := range map[string]string{
		"one": "on1e",
		"two": "tw2o",
		"three": "thr3ee",
		"four": "fo4ur",
		"five": "fi5ve",
		"six": "s6ix",
		"seven": "se7ven",
		"eight": "ei8ght",
		"nine": "ni9ne",
	} {
		s = strings.ReplaceAll(s, o, r)
	}

	return s
}

func ferror(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dat, err := os.ReadFile("input.txt")
	ferror(err)

	lines := strings.Split(string(dat), "\n")

	var total int64
	for n, r := range lines {
		l := sub(r)
		cal := ""
		for idx := 0; idx < len(l); idx++ {
			if l[idx] >= '0' && l[idx] <= '9' {
				cal = string(append([]byte(cal), l[idx]))
				break
			}
		}

		for idx := len(l) - 1; idx >= 0; idx-- {
			if l[idx] >= '0' && l[idx] <= '9' {
				cal = string(append([]byte(cal), l[idx]))

				break
			}
		}
		t, _ := strconv.Atoi(cal)
		fmt.Println(n, l, t)
		total += int64(t)
	}

	fmt.Println(total)
}

