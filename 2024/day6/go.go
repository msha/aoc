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

func deepCopy(original [][]rune) [][]rune {
	// Create a new slice with the same number of rows
	copy := make([][]rune, len(original))

	// Iterate through each row in the original slice
	for i := range original {
		// Create a new slice for each row
		copy[i] = make([]rune, len(original[i]))

		// Copy the contents of the row using a loop
		for j := range original[i] {
			copy[i][j] = original[i][j]
		}
	}

	return copy
}

type cords struct {
	x int
	y int
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
	var checklist []cords
	var locationx, locationy, startx, starty int
	linen := 0
	for scanner.Scan() {
		var line []rune
		for i, char := range scanner.Text() {
			line = append(line, char)
			if char == '^' {
				dir = 0
				startx, locationx = i, i
				starty, locationy = linen, linen
			}
		}
		area = append(area, line)
		linen += 1
	}
	total := 0
	nicemap := deepCopy(area)

	//part 1
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
			if area[locationy][locationx] != '^' {
				checklist = append(checklist, cords{x: locationx, y: locationy})
			}
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

	//part 2

	total2 := 0
	for _, cord := range checklist {

		locationx = startx
		locationy = starty
		dir = 0
		area = deepCopy(nicemap)
		area[cord.y][cord.x] = 'O'
		steps := 0
		for {
			if locationx >= len(area[0]) || locationy >= len(area) || locationx < 0 || locationy < 0 {

				break
			}
			if area[locationy][locationx] == '#' || area[locationy][locationx] == 'O' {
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
			if steps > 9999 {
				total2 += 1

				break
			}
			steps += 1
		}
	}

	fmt.Println("part1:", total)
	fmt.Println("part2:", total2)
}
