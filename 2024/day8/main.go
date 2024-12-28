package main

import (
	"bufio"
	"fmt"
	"os"
)

type cord struct {
	y int
	x int
	t rune
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var amap [][]rune
	var satellites []cord
	types := make(map[rune]rune)

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		var line []rune
		row := scanner.Text()
		for x, char := range row {
			line = append(line, char)
			if char != '.' {
				types[char] = char
				satellite := cord{y: y, x: x, t: char}
				satellites = append(satellites, satellite)
			}
		}
		amap = append(amap, line)
		y++
	}
	total := 0
	for _, t := range satellites {
		for _, s := range satellites {
			if s.t == t.t && (s.x != t.x && s.y != t.y) {
				infy := t.y + (t.y - s.y)
				infx := t.x + (t.x - s.x)
				inmapy := infy >= 0 && infy < len(amap)
				inmapx := false
				if inmapy {
					inmapx = infx >= 0 && infx < len(amap[infy])
				}
				if inmapy && inmapx {
					if amap[infy][infx] != '#' {
						amap[infy][infx] = '#'
						total += 1
					}

				}

			}

		}
	}

	for _, row := range amap {
		for _, r := range row {
			fmt.Print(string(r))
		}
		fmt.Println("")
	}
	fmt.Println(total)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
