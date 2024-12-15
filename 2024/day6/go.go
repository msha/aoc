package main

import (
	"bufio"
	"fmt"
	"os"
)

func printArea(area [][]rune) {
	for _, row := range area {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var area [][]rune // area[y][x]
	var dir int       // up = 0, right = 1, down = 2, left = 3
	var locationx, locationy int
	linen := 0
	for scanner.Scan() {
		var line []rune
		for i, char := range scanner.Text() {
			line = append(line, char)
			if char == '^' {
				dir = 0
				locationx = i
				locationy = linen
			}
		}
		area = append(area, line)
		linen += 1
	}
	total := 0

	for {
		if locationx >= len(area[0]) || locationy >= len(area) || locationx < 0 || locationy < 0 {
			break
		}
		if area[locationy][locationx] == '#' {
			dir += 1
			if dir == 4 {
				dir = 0
			}
			if dir == 0 {
				locationx += 1
			}
			if dir == 1 {
				locationy += 1
			}
			if dir == 2 {
				locationx -= 1
			}
			if dir == 3 {
				locationy -= 1
			}
		}
		if locationx >= len(area[0]) || locationy >= len(area) || locationx < 0 || locationy < 0 {
			break
		}
		if area[locationy][locationx] == '.' || area[locationy][locationx] == '^' {
			area[locationy][locationx] = 'X'
			total += 1
		}
		if dir == 0 {
			locationy -= 1
		}
		if dir == 1 {
			locationx += 1
		}
		if dir == 2 {
			locationy += 1
		}
		if dir == 3 {
			locationx -= 1
		}
	}

	fmt.Println(dir, locationx, locationy)
	printArea(area)

	fmt.Println(total)
}
