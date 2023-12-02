package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("q02/input.txt")
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mem := make(map[string]int)
		line := scanner.Text()
		s := strings.Split(line, ": ")
		processEvents(strings.Split(s[1], "; "), mem)
		pow := 1
		for _, count := range mem {
			pow *= count
		}
		total += pow
	}

	fmt.Println(total)
}

func processEvents(events []string, out map[string]int) {
	for _, event := range events {
		for _, action := range strings.Split(event, ", ") {
			tmp := strings.Split(action, " ")
			count, _ := strconv.Atoi(tmp[0])
			color := tmp[1]
			curr, ok := out[color]
			if !ok || count > curr {
				out[color] = count
			}
		}
	}
}
