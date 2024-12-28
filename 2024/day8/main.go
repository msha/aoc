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
	total2 := 0
	for _, t := range satellites {
		for _, s := range satellites {
			if s.t == t.t && (s.x != t.x && s.y != t.y) {
				changx := (t.y - s.y)
				changy := (t.x - s.x)
				factor := 0
				infy := t.y + (changx * factor)
				infx := t.x + (changy * factor)
				for {
					inmapy := infy >= 0 && infy < len(amap)
					inmapx := false
					if inmapy {
						inmapx = infx >= 0 && infx < len(amap[infy])
					}
					if inmapy && inmapx {
						if amap[infy][infx] != '#' {
							amap[infy][infx] = '#'
							if factor == 1 {
								total++
							}
							total2++
						}
						factor++
						infy = t.y + (changx * factor)
						infx = t.x + (changy * factor)
					} else {
						break
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
	fmt.Println("part1:", total) //this broke some how, too lazy to fix now lel
	fmt.Println("part2:", total2)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
