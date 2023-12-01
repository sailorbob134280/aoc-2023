package main

import (
	"bufio"
	"os"
	"strconv"
)

func numeric(c string) (int, bool) {
	if num, err := strconv.Atoi(c); err == nil {
		return num, true
	}

	return 0, false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		first, last := 0, 0
		first_found := false

		line := scanner.Text()
		for _, c := range line {
			if num, is_num := numeric(string(c)); is_num {
				last = num
				if !first_found {
					first = num
					first_found = true
				}
			}
		}

		sum += first*10 + last

	}

	print("Sum of calibrations: ", sum, "\n")
}
