package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Number struct {
	idx   int
	value int
}

func main() {
	numeric_strings := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers_found := []Number{}

		line := scanner.Text()
		for str, val := range numeric_strings {

			first_idx := strings.Index(line, str)

			if first_idx >= 0 {
				numbers_found = append(numbers_found, Number{first_idx, val})
			}

			last_idx := strings.LastIndex(line, str)

			if last_idx >= 0 {
				numbers_found = append(numbers_found, Number{last_idx, val})
			}
		}

		sort.SliceStable(numbers_found, func(i, j int) bool {
			return numbers_found[i].idx < numbers_found[j].idx
		})

		sum += numbers_found[0].value*10 + numbers_found[len(numbers_found)-1].value

	}

	print("Sum of calibrations: ", sum, "\n")
}
