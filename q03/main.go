package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, _ := os.Open("q03/input.txt")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	total := 0

	for y, line := range lines {
		for x, c := range line {
			if c == '*' {
				nums := exploreAdjacentPartNumbers(lines, y, x)
				if len(nums) == 2 {
					total += nums[0] * nums[1]
				}
			}
		}
	}

	fmt.Println(total)
}

func exploreAdjacentPartNumbers(lines []string, y int, x int) []int {
	explored := make([][]bool, len(lines))
	for i := range explored {
		explored[i] = make([]bool, len(lines[0]))
	}
	var nums []int
	findPart(lines, &explored, y-1, x-1, &nums)
	findPart(lines, &explored, y-1, x, &nums)
	findPart(lines, &explored, y-1, x+1, &nums)
	findPart(lines, &explored, y, x-1, &nums)
	findPart(lines, &explored, y, x+1, &nums)
	findPart(lines, &explored, y+1, x-1, &nums)
	findPart(lines, &explored, y+1, x, &nums)
	findPart(lines, &explored, y+1, x+1, &nums)

	return nums
}

func findPart(lines []string, explored *[][]bool, y int, x int, out *[]int) {
	ex := *explored
	if ex[y][x] {
		return
	}
	ex[y][x] = true
	if y >= len(lines) || y < 0 {
		return
	}
	if x >= len(lines) || y < 0 {
		return
	}
	line := lines[y]
	if unicode.IsDigit(rune(line[x])) {
		sX, eX := findNumberBoundaries(line, x)
		for i := sX; i <= eX; i++ {
			ex[y][i] = true
		}
		*out = append(*out, parseNum(line, sX, eX))
	}
	*explored = ex
}

func findNumberBoundaries(line string, x int) (int, int) {
	var sX = x
	var eX = x
	for sX >= 0 && unicode.IsDigit(rune(line[sX])) {
		sX--
	}
	sX++
	for eX < len(line) && unicode.IsDigit(rune(line[eX])) {
		eX++
	}
	eX--
	return sX, eX
}

func parseNum(line string, sX, eX int) int {
	val, _ := strconv.Atoi(line[sX : eX+1])
	return val
}
