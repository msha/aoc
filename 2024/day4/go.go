package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Collect forward diagonals (Top-Left to Bottom-Right)
func collectForwardDiagonals(matrix []string) []string {
	n := len(matrix)
	if n == 0 {
		return nil
	}
	m := len(matrix[0])

	var diagonals []string

	for d := 0; d < n+m-1; d++ {
		diagonal := ""
		for i := 0; i < n; i++ {
			j := d - i
			if j >= 0 && j < m {
				diagonal += string(matrix[i][j])
			}
		}
		diagonals = append(diagonals, diagonal)
	}
	return diagonals
}

// Collect backward diagonals (Top-Right to Bottom-Left)
func collectBackwardDiagonals(matrix []string) []string {
	n := len(matrix)
	if n == 0 {
		return nil
	}
	m := len(matrix[0])

	var diagonals []string

	for d := 0; d < n+m-1; d++ {
		diagonal := ""
		for i := 0; i < n; i++ {
			j := m - 1 - d + i
			if j >= 0 && j < m {
				diagonal += string(matrix[i][j])
			}
		}
		diagonals = append(diagonals, diagonal)
	}
	return diagonals
}

func rotateClockwise(matrix []string) []string {
	n := len(matrix)
	if n == 0 {
		return nil
	}
	m := len(matrix[0])

	rotated := make([]string, m)
	for j := 0; j < m; j++ {
		newRow := ""
		for i := n - 1; i >= 0; i-- {
			newRow += string(matrix[i][j])
		}
		rotated[j] = newRow
	}
	return rotated
}

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func countXMAS(s string) int {
	re := regexp.MustCompile(`XMAS`)
	matches := re.FindAllString(s, -1)
	return len(matches)
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrix []string
	var total int
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	for _, line := range matrix {
		total += countXMAS(line)
		total += countXMAS(reverseString(line))
	}

	for _, line := range rotateClockwise(matrix) {
		total += countXMAS(reverseString(line))
		total += countXMAS(line)
	}

	forwardDiagonals := collectForwardDiagonals(matrix)
	reversedForwardDiagonals := make([]string, len(forwardDiagonals))
	for i, diag := range forwardDiagonals {
		reversedForwardDiagonals[i] = reverseString(diag)
	}

	backwardDiagonals := collectBackwardDiagonals(matrix)
	reversedBackwardDiagonals := make([]string, len(backwardDiagonals))
	for i, diag := range backwardDiagonals {
		reversedBackwardDiagonals[i] = reverseString(diag)
	}

	for _, diag := range forwardDiagonals {
		total += countXMAS(diag)
		total += countXMAS(reverseString(diag))
	}

	for _, diag := range backwardDiagonals {
		total += countXMAS(diag)
		total += countXMAS(reverseString(diag))
	}

	fmt.Println("\nTotal:", total)
}
