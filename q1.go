package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"unicode"
)

var words = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	file, _ := os.Open("q1.txt")
	defer file.Close()

	start := time.Now()
	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var first int
		var last int

		positions := make([]int, 9)
		for _, c := range line {
			digit := parseDigit(c, positions, false)
			if digit != -1 {
				first = digit
				break
			}
		}

		positions = make([]int, 9)
		for i := len(line) - 1; i >= 0; i-- {
			digit := parseDigit(rune(line[i]), positions, true)
			if digit != -1 {
				last = digit
				break
			}
		}

		total += 10 * first
		total += last
	}

	fmt.Println(total)

	elapsed := time.Since(start)
	fmt.Printf("Took %s", elapsed)
}

func parseDigit(c rune, positions []int, backwards bool) int {
	if unicode.IsDigit(c) {
		return int(c - '0')
	}

	for i, pos := range positions {
		expectedRune := runeAt(pos, i, backwards)
		if expectedRune == c {
			positions[i]++
		} else if runeAt(0, i, backwards) == c {
			positions[i] = 1
		} else {
			positions[i] = 0
		}
		if positions[i] == len(words[i]) {
			return i + 1
		}
	}

	return -1
}

func runeAt(pos, wordIndex int, backwards bool) rune {
	word := words[wordIndex]
	if backwards {
		pos = len(word) - pos - 1
	}
	return rune(word[pos])
}
