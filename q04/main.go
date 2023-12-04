package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("q04/input.txt")
	defer file.Close()

	var cards []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cards = append(cards, scanner.Text())
	}

	winsPerCard := make([]int, len(cards))

	for _, card := range cards {
		s := strings.Split(card, ": ")
		cardIndex, _ := strconv.Atoi(strings.Fields(s[0])[1])
		cardIndex--

		s = strings.Split(s[1], " | ")
		winningNums := strings.Fields(s[0])
		nums := strings.Fields(s[1])

		wins := 0
		for _, num := range nums {
			if contains(winningNums, num) {
				wins++
			}
		}
		winsPerCard[cardIndex] = wins
	}

	totalCards := 0
	for i := 0; i < len(cards); i++ {
		totalCards += countCards(i, winsPerCard)
	}

	fmt.Println(totalCards)
}

func countCards(cardIndex int, winsPerCard []int) int {
	total := 1
	for i := 1; i <= winsPerCard[cardIndex]; i++ {
		total += countCards(cardIndex+i, winsPerCard)
	}
	return total
}

func contains(winningNums []string, num string) bool {
	for _, wN := range winningNums {
		if num == wN {
			return true
		}
	}
	return false
}
