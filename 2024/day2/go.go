package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		numbers := make([]int, len(words))
		for i, word := range words {
			num, err := strconv.Atoi(word)
			if err != nil {
				return
			}
			numbers[i] = num
		}
		fail := 0
		if numbers[0] == numbers[1] {
			fail += 1
		}
		dir := numbers[fail] < numbers[fail+1]

		value := numbers[fail]

		for _, w := range numbers[fail+1:] {
			diff := value - w
			if dir {
				diff = -diff
			}
			if diff < 1 || diff > 3 {
				fail += 1
			} else {
				value = w
			}
		}
		if fail < 1 {
			total += 1
		}
	}
	println("total:", total)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
