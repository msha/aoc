package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func containsMAS(s string) bool {
	return strings.Count(s, "M") == 2 && strings.Count(s, "S") == 2 && (strings.Count(s, "MM") == 1 || strings.Count(s, "SS") == 1)
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrix [][]string
	var total int
	for scanner.Scan() {
		text := scanner.Text()
		for i, char := range text {
			if len(matrix) <= i {
				matrix = append(matrix, []string{})
			}
			matrix[i] = append(matrix[i], string(char))
		}
	}
	for y, row := range matrix[1 : len(matrix)-1] {

		for x, char := range row[1 : len(row)-1] {
			if char == "A" {
				word := matrix[y][x] + matrix[y][x+2] + matrix[y+2][x+2] + matrix[y+2][x]
				if containsMAS(word) {
					fmt.Println(y, x, word)
					total += 1
				}
			}
		}
	}

	fmt.Println("\nTotal:", total)
}
